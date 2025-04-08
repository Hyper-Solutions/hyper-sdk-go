package internal

import (
	"bytes"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"sync"
)

var zstdEncoderPool = sync.Pool{
	New: func() interface{} {
		encoder, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
		if err != nil {
			panic(fmt.Sprintf("failed to create zstd encoder: %v", err))
		}
		return encoder
	},
}

func CompressZstd(data []byte) ([]byte, error) {
	// Get an encoder from the pool
	encoderInterface := zstdEncoderPool.Get()
	encoder := encoderInterface.(*zstd.Encoder)

	// Make sure we return the encoder to the pool when we're done
	defer zstdEncoderPool.Put(encoder)

	// Create a buffer to hold the compressed data
	var buf bytes.Buffer

	// Reset the encoder with our buffer
	encoder.Reset(&buf)

	// Write the data to the encoder
	if _, err := encoder.Write(data); err != nil {
		return nil, err
	}

	// Close the encoder to flush any remaining data
	if err := encoder.Close(); err != nil {
		return nil, err
	}

	// Return the compressed data
	return buf.Bytes(), nil
}
