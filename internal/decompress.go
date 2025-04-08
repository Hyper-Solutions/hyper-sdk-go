package internal

import (
	"bytes"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zlib"
	"github.com/klauspost/compress/zstd"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Decoder pools
var (
	zstdDecoderPool = sync.Pool{
		New: func() interface{} {
			decoder, err := zstd.NewReader(nil)
			if err != nil {
				panic(err)
			}
			return decoder
		},
	}

	gzipDecoderPool = sync.Pool{
		New: func() interface{} {
			return new(gzip.Reader)
		},
	}
)

// DecompressResponse decompresses the response body based on Content-Encoding header
func DecompressResponse(resp *http.Response) ([]byte, error) {
	if resp == nil || resp.Body == nil {
		return nil, nil
	}

	encoding := strings.ToLower(resp.Header.Get("Content-Encoding"))
	if encoding == "" {
		return io.ReadAll(resp.Body)
	}

	switch {
	case strings.Contains(encoding, "gzip"):
		return handleGzip(resp.Body)
	case strings.Contains(encoding, "zstd"):
		return handleZstd(resp.Body)
	case strings.Contains(encoding, "br"):
		return handleBrotli(resp.Body)
	case strings.Contains(encoding, "deflate"):
		return handleDeflate(resp.Body)
	default:
		// Unknown encoding, try to read as-is
		return io.ReadAll(resp.Body)
	}
}

func handleGzip(body io.ReadCloser) ([]byte, error) {
	reader := gzipDecoderPool.Get().(*gzip.Reader)
	defer gzipDecoderPool.Put(reader)

	if err := reader.Reset(body); err != nil {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

func handleZstd(body io.ReadCloser) ([]byte, error) {
	decoder := zstdDecoderPool.Get().(*zstd.Decoder)
	defer zstdDecoderPool.Put(decoder)

	decoder.Reset(body)
	return io.ReadAll(decoder)
}

func handleBrotli(body io.ReadCloser) ([]byte, error) {
	brReader := brotli.NewReader(body)
	return io.ReadAll(brReader)
}

func handleDeflate(body io.ReadCloser) ([]byte, error) {
	// Read the entire compressed body first
	compressedData, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	// First try with zlib (RFC 1950) which is the correct format for HTTP "deflate"
	zlibReader, err := zlib.NewReader(bytes.NewReader(compressedData))
	if err == nil {
		defer zlibReader.Close()
		return io.ReadAll(zlibReader)
	}

	// If zlib fails, try raw deflate as a fallback (RFC 1951)
	// Some servers incorrectly send raw deflate data without the zlib wrapper
	rawReader := flate.NewReader(bytes.NewReader(compressedData))
	defer rawReader.Close()
	return io.ReadAll(rawReader)
}
