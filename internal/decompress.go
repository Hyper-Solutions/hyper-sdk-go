package internal

import (
	"bytes"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/gzip"
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
