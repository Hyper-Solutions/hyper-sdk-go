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
	encoderInterface := zstdEncoderPool.Get()
	encoder := encoderInterface.(*zstd.Encoder)
	defer zstdEncoderPool.Put(encoder)

	var buf bytes.Buffer

	encoder.Reset(&buf)

	if _, err := encoder.Write(data); err != nil {
		return nil, err
	}

	if err := encoder.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
