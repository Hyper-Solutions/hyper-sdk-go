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

	var (
		body []byte
		err  error
	)

	switch {
	case strings.Contains(encoding, "gzip"):
		body, err = handleGzip(resp.Body)
	case strings.Contains(encoding, "zstd"):
		body, err = handleZstd(resp.Body)
	case strings.Contains(encoding, "br"):
		body, err = handleBrotli(resp.Body)
	case strings.Contains(encoding, "deflate"):
		body, err = handleDeflate(resp.Body)
	default:
		// Unknown encoding, try to read as-is
		return io.ReadAll(resp.Body)
	}

	resp.Body.Close()

	resp.Body = io.NopCloser(bytes.NewBuffer(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Del("Content-Encoding")

	return body, err
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
	// The "deflate" content encoding is supposed to be the zlib format (RFC 1950)
	// which wraps the deflate format (RFC 1951) with a header and checksum
	reader, err := zlib.NewReader(body)
	if err != nil {
		// If zlib fails, try raw deflate as a fallback
		// (some servers incorrectly use raw deflate)
		bodyBytes, readErr := io.ReadAll(body)
		if readErr != nil {
			return nil, readErr
		}

		rawReader := flate.NewReader(bytes.NewReader(bodyBytes))
		defer rawReader.Close()
		return io.ReadAll(rawReader)
	}

	defer reader.Close()
	return io.ReadAll(reader)
}
