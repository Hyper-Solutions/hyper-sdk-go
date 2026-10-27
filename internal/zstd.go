package internal

import (
	"runtime"

	"github.com/klauspost/compress/zstd"
)

// zstdEncoder compresses every request with EncodeAll, which is safe for
// concurrent use and keeps its block encoders for the life of the process, so
// a request does not pay for setting up an encoder.
var zstdEncoder = func() *zstd.Encoder {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderConcurrency(min(runtime.GOMAXPROCS(0), 8)), zstd.WithWindowSize(1<<18))
	if err != nil {
		panic(err)
	}
	return encoder
}()

// CompressZstd compresses data as one zstd frame.
func CompressZstd(data []byte) ([]byte, error) {
	return zstdEncoder.EncodeAll(data, make([]byte, 0, len(data)/4)), nil
}
