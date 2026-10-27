package internal

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/klauspost/compress/gzip"
)

var gzipWriterPool = sync.Pool{New: func() any { return gzip.NewWriter(nil) }}

// CompressGzip compresses data using gzip
func CompressGzip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, len(data)/4))

	gzWriter := gzipWriterPool.Get().(*gzip.Writer)
	defer gzipWriterPool.Put(gzWriter)
	gzWriter.Reset(buf)

	if _, err := gzWriter.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write to gzip writer: %w", err)
	}

	if err := gzWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}
