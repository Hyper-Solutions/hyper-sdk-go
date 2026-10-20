package hyper

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mailru/easyjson"
)

// stubTransport answers every request with body and records the last request
// body, so the Generate* methods run end to end without the network.
type stubTransport struct {
	body     string
	lastBody string
}

func (t *stubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		b, _ := io.ReadAll(req.Body)
		t.lastBody = string(b)
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(t.body)),
		Request:    req,
	}, nil
}

func stubSession(body string) (*Session, *stubTransport) {
	tr := &stubTransport{body: body}
	return NewSession("test-key").WithClient(&http.Client{Transport: tr}), tr
}

// The API's "headers" object reaches the caller as ClientHints, every key
// intact, on each product that returns it.
func TestClientHintsReachTheResult(t *testing.T) {
	const body = `{"payload":"p","context":"c","headers":{"sec-ch-ua":"\"Chromium\";v=\"154\"","sec-ch-prefers-color-scheme":"dark","rtt":"150"}}`
	ctx := context.Background()

	check := func(t *testing.T, hints ClientHints) {
		t.Helper()
		want := map[string]string{
			"sec-ch-ua":                   `"Chromium";v="154"`,
			"sec-ch-prefers-color-scheme": "dark",
			"rtt":                         "150",
		}
		if len(hints) != len(want) {
			t.Fatalf("got %d hints, want %d: %v", len(hints), len(want), hints)
		}
		for k, v := range want {
			if hints[k] != v {
				t.Errorf("%s = %q, want %q", k, hints[k], v)
			}
		}
	}

	t.Run("akamai sensor", func(t *testing.T) {
		s, _ := stubSession(body)
		res, err := s.GenerateSensorData(ctx, &SensorInput{})
		if err != nil {
			t.Fatal(err)
		}
		if res.Payload != "p" || res.Context != "c" {
			t.Errorf("payload/context = %q/%q", res.Payload, res.Context)
		}
		check(t, res.ClientHints)
	})
	t.Run("akamai sbsd", func(t *testing.T) {
		s, _ := stubSession(body)
		res, err := s.GenerateSbsdData(ctx, &SbsdInput{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, res.ClientHints)
	})
	t.Run("incapsula reese84", func(t *testing.T) {
		s, _ := stubSession(body)
		res, err := s.GenerateReese84Sensor(ctx, &ReeseInput{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, res.ClientHints)
	})
	t.Run("datadome", func(t *testing.T) {
		s, _ := stubSession(body)
		slider, err := s.GenerateDataDomeSlider(ctx, &DataDomeSliderInput{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, slider.ClientHints)
		inter, err := s.GenerateDataDomeInterstitial(ctx, &DataDomeInterstitialInput{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, inter.ClientHints)
		tags, err := s.GenerateDataDomeTags(ctx, &DataDomeTagsInput{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, tags.ClientHints)
	})
}

// Kasada keeps its x-kpsdk-* headers under "headers" and returns the hints
// under "clientHints"; the two must not be mixed up.
func TestKasadaPayloadSplitsHeadersAndClientHints(t *testing.T) {
	payload := base64.StdEncoding.EncodeToString([]byte("tl-body"))
	s, _ := stubSession(`{"payload":"` + payload + `","headers":{"x-kpsdk-ct":"ct","x-kpsdk-v":"j-1.0.0"},"clientHints":{"rtt":"100","sec-ch-ua-wow64":"?0"}}`)

	res, err := s.GenerateKasadaPayload(context.Background(), &KasadaPayloadInput{})
	if err != nil {
		t.Fatal(err)
	}
	if string(res.Payload) != "tl-body" {
		t.Errorf("payload = %q", res.Payload)
	}
	if res.Headers.XKpsdkCt != "ct" || res.Headers.XKpsdkV != "j-1.0.0" {
		t.Errorf("kpsdk headers = %+v", res.Headers)
	}
	if res.ClientHints["rtt"] != "100" || res.ClientHints["sec-ch-ua-wow64"] != "?0" || len(res.ClientHints) != 2 {
		t.Errorf("client hints = %v", res.ClientHints)
	}
	if _, ok := res.ClientHints["x-kpsdk-ct"]; ok {
		t.Error("an x-kpsdk header leaked into ClientHints")
	}
}

// No hints from the API (Safari, remote-browser payloads, older deployments)
// means a nil map, not an error.
func TestMissingClientHintsAreNil(t *testing.T) {
	s, _ := stubSession(`{"payload":"p"}`)
	res, err := s.GenerateReese84Sensor(context.Background(), &ReeseInput{})
	if err != nil {
		t.Fatal(err)
	}
	if res.ClientHints != nil {
		t.Errorf("ClientHints = %v, want nil", res.ClientHints)
	}
}

// Kasada strict and SBSD scriptUrl are optional: absent from the body unless
// set, and sent when set (models_easyjson.go is generated, so a field missing
// from the codec would silently never reach the API).
func TestStrictAndScriptUrlOnTheWire(t *testing.T) {
	cases := []struct {
		name  string
		input easyjson.Marshaler
		key   string
		want  bool
	}{
		{"kasada strict unset", &KasadaPayloadInput{}, `"strict"`, false},
		{"kasada strict set", &KasadaPayloadInput{Strict: true}, `"strict":true`, true},
		{"sbsd scriptUrl unset", &SbsdInput{}, `"scriptUrl"`, false},
		{"sbsd scriptUrl set", &SbsdInput{ScriptUrl: "https://example.com/abc.js?v=1"}, `"scriptUrl":"https://example.com/abc.js?v=1"`, true},
	}
	for _, tc := range cases {
		raw, err := easyjson.Marshal(tc.input)
		if err != nil {
			t.Fatal(err)
		}
		if got := strings.Contains(string(raw), tc.key); got != tc.want {
			t.Errorf("%s: %s in %s = %v, want %v", tc.name, tc.key, raw, got, tc.want)
		}
	}
}
