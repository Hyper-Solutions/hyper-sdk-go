package hyper

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// fakeScriptAPI answers like the API: it stores the scripts it is sent, solves
// a request that names a stored script by hash, and answers 422
// script_required for one it does not know.
type fakeScriptAPI struct {
	mu       sync.Mutex
	known    map[string]bool
	requests []map[string]string
	failWith string // answer every script request with this error
}

func (f *fakeScriptAPI) RoundTrip(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	f.serve(rec, req)
	return rec.Result(), nil
}

func (f *fakeScriptAPI) serve(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	var in map[string]string
	_ = json.Unmarshal(body, &in)

	f.mu.Lock()
	defer f.mu.Unlock()
	f.requests = append(f.requests, map[string]string{"script": in["script"], "scriptHash": in["scriptHash"]})
	switch {
	case in["script"] != "" && f.failWith != "":
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"` + f.failWith + `"}`))
	case in["script"] != "":
		f.known[ScriptHash(in["script"])] = true
		_, _ = w.Write([]byte(`{"payload":"solved"}`))
	case f.known[in["scriptHash"]]:
		_, _ = w.Write([]byte(`{"payload":"solved"}`))
	default:
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"error":"script_required"}`))
	}
}

// sent returns what each request since the last call carried: "script",
// "hash" or "".
func (f *fakeScriptAPI) sent() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	var out []string
	for _, r := range f.requests {
		switch {
		case r["script"] != "" && r["scriptHash"] != "":
			out = append(out, "both")
		case r["script"] != "":
			out = append(out, "script")
		case r["scriptHash"] != "":
			out = append(out, "hash")
		default:
			out = append(out, "")
		}
	}
	f.requests = nil
	return out
}

func newFakeScriptSession() (*Session, *fakeScriptAPI) {
	api := &fakeScriptAPI{known: map[string]bool{}}
	return NewSession("key").WithClient(&http.Client{Transport: api}), api
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestScriptHashIsSha256Hex(t *testing.T) {
	// echo -n "abc" | sha256sum
	const want = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if got := ScriptHash("abc"); got != want {
		t.Fatalf("ScriptHash(abc) = %s, want %s", got, want)
	}
}

func TestDataDomeSendsScriptHashFirst(t *testing.T) {
	ctx := context.Background()
	const bundle = "/* DataDome v1.36.0 */ !function(){}();"

	for name, call := range map[string]func(*Session) error{
		"interstitial": func(s *Session) error {
			_, err := s.GenerateDataDomeInterstitial(ctx, &DataDomeInterstitialInput{Html: "<html/>", Script: bundle})
			return err
		},
		"slider": func(s *Session) error {
			_, err := s.GenerateDataDomeSlider(ctx, &DataDomeSliderInput{Html: "<html/>", Script: bundle})
			return err
		},
	} {
		t.Run(name, func(t *testing.T) {
			s, api := newFakeScriptSession()

			steps := []struct {
				name string
				do   func()
				want []string
			}{
				{"an unseen script is sent after the 422", nil, []string{"hash", "script"}},
				{"then only the hash is sent", nil, []string{"hash"}},
				{"a script another session sent is never uploaded", func() { s = NewSession("key").WithClient(&http.Client{Transport: api}) }, []string{"hash"}},
				{"an evicted script is sent again", func() { delete(api.known, ScriptHash(bundle)) }, []string{"hash", "script"}},
			}
			for _, step := range steps {
				if step.do != nil {
					step.do()
				}
				if err := call(s); err != nil {
					t.Fatalf("%s: %v", step.name, err)
				}
				if got := api.sent(); !equal(got, step.want) {
					t.Fatalf("%s: sent %v, want %v", step.name, got, step.want)
				}
			}
		})
	}
}

// TestDataDomeErrorAfterTheScript checks that an error answer to the request
// that carried the script is returned, and the input is left unchanged.
func TestDataDomeErrorAfterTheScript(t *testing.T) {
	ctx := context.Background()
	s, api := newFakeScriptSession()
	api.failWith = "failed to generate interstitial payload"
	input := &DataDomeInterstitialInput{Html: "<html/>", Script: "bundle"}

	if _, err := s.GenerateDataDomeInterstitial(ctx, input); err == nil {
		t.Fatal("expected the API error")
	}
	api.failWith = ""
	if _, err := s.GenerateDataDomeInterstitial(ctx, input); err != nil {
		t.Fatal(err)
	}
	if got := api.sent(); !equal(got, []string{"hash", "script", "hash", "script"}) {
		t.Fatalf("sent %v, want hash then script both times", got)
	}
	if input.ScriptHash != "" {
		t.Fatalf("the caller's input was modified: ScriptHash = %q", input.ScriptHash)
	}
}

func TestDataDomeHashWithoutScript(t *testing.T) {
	ctx := context.Background()
	s, api := newFakeScriptSession()
	_, err := s.GenerateDataDomeSlider(ctx, &DataDomeSliderInput{Html: "<html/>", ScriptHash: ScriptHash("never sent")})
	if !errors.Is(err, ErrScriptRequired) {
		t.Fatalf("error = %v, want ErrScriptRequired", err)
	}
	if got := api.sent(); !equal(got, []string{"hash"}) {
		t.Fatalf("sent %v, want one hash-only request", got)
	}
}
