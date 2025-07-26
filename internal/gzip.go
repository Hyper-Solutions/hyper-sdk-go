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
	defer gzWriter.Close()

	if _, err := gzWriter.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write to gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}
