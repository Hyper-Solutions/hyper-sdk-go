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
			// Only DecodeAll is used, which runs on the caller's goroutine.
			decoder, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
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

// DecompressResponse decompresses the response body based on Content-Encoding
// header. The body can be held in pooled buffers: release returns them, and
// the bytes must not be used after it is called.
func DecompressResponse(resp *http.Response) (body []byte, release func(), err error) {
	release = func() {}
	if resp == nil || resp.Body == nil {
		return nil, release, nil
	}

	encoding := strings.ToLower(resp.Header.Get("Content-Encoding"))
	switch {
	case strings.Contains(encoding, "gzip"):
		dst := getBuffer(-1)
		body, err = handleGzip(resp.Body, dst)
		return body, func() { putBuffer(dst) }, err
	case strings.Contains(encoding, "zstd"):
		// The compressed body is read once into a pooled buffer and decoded
		// with DecodeAll into another.
		src := getBuffer(resp.ContentLength)
		defer putBuffer(src)
		compressed, err := readInto(src, resp.Body)
		if err != nil {
			return nil, release, err
		}
		dst := getBuffer(-1)
		body, err = handleZstd(compressed, dst)
		return body, func() { putBuffer(dst) }, err
	case strings.Contains(encoding, "br"):
		body, err = handleBrotli(resp.Body)
		return body, release, err
	case strings.Contains(encoding, "deflate"):
		body, err = handleDeflate(resp.Body)
		return body, release, err
	default:
		// No or unknown encoding, read as-is
		dst := getBuffer(resp.ContentLength)
		body, err = readInto(dst, resp.Body)
		return body, func() { putBuffer(dst) }, err
	}
}

// maxPooledBuffer keeps a rare huge body from pinning its buffer in the pool.
const maxPooledBuffer = 8 << 20

var bufferPool = sync.Pool{New: func() any { return new([]byte) }}

// getBuffer returns a pooled buffer with room for at least size bytes.
func getBuffer(size int64) *[]byte {
	buf := bufferPool.Get().(*[]byte)
	// The extra room lets the read that reports EOF land without growing.
	if size > 0 && size <= maxPooledBuffer && int64(cap(*buf)) < size+bytes.MinRead {
		*buf = make([]byte, 0, size+bytes.MinRead)
	}
	*buf = (*buf)[:0]
	return buf
}

func putBuffer(buf *[]byte) {
	if cap(*buf) > maxPooledBuffer {
		return
	}
	bufferPool.Put(buf)
}

// readInto reads r to EOF into buf, growing it as needed, and returns the
// bytes read. buf keeps the grown slice for the pool.
func readInto(buf *[]byte, r io.Reader) ([]byte, error) {
	b := (*buf)[:0]
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			*buf = b
			if err == io.EOF {
				return b, nil
			}
			return b, err
		}
	}
}

func handleGzip(body io.Reader, dst *[]byte) ([]byte, error) {
	reader := gzipDecoderPool.Get().(*gzip.Reader)
	defer gzipDecoderPool.Put(reader)

	if err := reader.Reset(body); err != nil {
		return nil, err
	}
	defer reader.Close()

	return readInto(dst, reader)
}

func handleZstd(compressed []byte, dst *[]byte) ([]byte, error) {
	decoder := zstdDecoderPool.Get().(*zstd.Decoder)
	defer zstdDecoderPool.Put(decoder)

	body, err := decoder.DecodeAll(compressed, (*dst)[:0])
	if cap(body) <= maxPooledBuffer {
		*dst = body[:0]
	}
	return body, err
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
