package hyper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/klauspost/compress/zstd"
)

// benchTransport reads the request body and answers with a fixed
// zstd-encoded response, so a benchmark measures the SDK's own encoding,
// decoding and JSON work.
type benchTransport struct{ response []byte }

func (b benchTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_, _ = io.Copy(io.Discard, req.Body)
	_ = req.Body.Close()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Encoding": []string{"zstd"}},
		Body:       io.NopCloser(bytes.NewReader(b.response)),
		Request:    req,
	}, nil
}

// benchScript is about 600 KB of script-like text, the size of a challenge
// script.
func benchScript() string {
	var sb strings.Builder
	for i := 0; sb.Len() < 600<<10; i++ {
		fmt.Fprintf(&sb, "var _0x%04x=function(a,b){return a[%d]^b.charCodeAt(%d)};", i, i%97, i%13)
	}
	return sb.String()
}

func benchResponse(tb testing.TB) []byte {
	enc, err := zstd.NewWriter(nil)
	if err != nil {
		tb.Fatal(err)
	}
	body := `{"payload":"` + strings.Repeat("a1b2c3d4", 10<<10/8) + `","headers":{"sec-ch-ua":"x"}}`
	return enc.EncodeAll([]byte(body), nil)
}

// BenchmarkScriptRequestRoundTrip sends a request carrying a challenge script
// (zstd encoded) and decodes a zstd response.
func BenchmarkScriptRequestRoundTrip(b *testing.B) {
	s := NewSession("key").WithClient(&http.Client{Transport: benchTransport{response: benchResponse(b)}})
	input := &DataDomeInterstitialInput{UserAgent: "UA", Html: benchScript(), DeviceLink: "https://example.com/"}
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sendRequest[*DataDomeInterstitialInput, *apiResponse](ctx, s, "https://example.com/interstitial", input); err != nil {
			b.Fatal(err)
		}
	}
}
