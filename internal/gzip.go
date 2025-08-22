package internal

import (
	"bytes"
	"fmt"

	"github.com/klauspost/compress/gzip"
)

// CompressGzip compresses data using gzip
func CompressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	gzWriter := gzip.NewWriter(&buf)

	if _, err := gzWriter.Write(data); err != nil {
		gzWriter.Close()
		return nil, fmt.Errorf("failed to write to gzip writer: %w", err)
	}

	if err := gzWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}
