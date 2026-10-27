package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

// TestRoundTripPooled compresses and decompresses bodies of several sizes
// with every request encoding, from many goroutines, so pooled encoders,
// decoders and buffers are reused across requests. Run it with -race.
func TestRoundTripPooled(t *testing.T) {
	encodings := map[string]func([]byte) ([]byte, error){
		"zstd":     CompressZstd,
		"gzip":     CompressGzip,
		"identity": func(b []byte) ([]byte, error) { return append([]byte(nil), b...), nil },
	}
	var bodies [][]byte
	for _, n := range []int{0, 1, 999, 64 << 10, 700 << 10} {
		var sb strings.Builder
		for i := 0; sb.Len() < n; i++ {
			fmt.Fprintf(&sb, "{\"k%d\":\"v%d\"},", i, i*7)
		}
		bodies = append(bodies, []byte(sb.String()[:n]))
	}

	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				for name, compress := range encodings {
					want := bodies[(g+i)%len(bodies)]
					compressed, err := compress(want)
					if err != nil {
						t.Errorf("%s compress: %v", name, err)
						return
					}
					resp := &http.Response{Header: http.Header{}, Body: io.NopCloser(bytes.NewReader(compressed)), ContentLength: int64(len(compressed))}
					if name != "identity" {
						resp.Header.Set("Content-Encoding", name)
					}
					got, release, err := DecompressResponse(resp)
					if err != nil {
						t.Errorf("%s decompress: %v", name, err)
						return
					}
					if !bytes.Equal(got, want) {
						t.Errorf("%s round trip of %d bytes returned %d bytes", name, len(want), len(got))
					}
					release()
				}
			}
		}(g)
	}
	wg.Wait()
}
