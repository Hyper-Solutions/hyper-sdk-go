package hyper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// usageServer answers every request with status and body, and sends the
// session's requests to it in place of the real API.
func usageServer(t *testing.T, status int, body string) (*Session, *[]*http.Request) {
	t.Helper()
	var seen []*http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Clone(context.Background()))
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)

	target, _ := url.Parse(srv.URL)
	client := &http.Client{Transport: rewriteTransport{target: target}}
	return NewSession("test-key").WithClient(client), &seen
}

type rewriteTransport struct{ target *url.URL }

func (rt rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("x-original-url", req.URL.String())
	req.URL.Scheme, req.URL.Host = rt.target.Scheme, rt.target.Host
	return http.DefaultTransport.RoundTrip(req)
}

func TestGetUsage(t *testing.T) {
	s, seen := usageServer(t, http.StatusOK, `{"balanceEuro":12.5,"plans":[{"expiresAt":"2027-07-31T00:00:00Z",`+
		`"requestsUsed":44235934,"requestsQuota":9999999999,"requestsRemaining":9955764065,"product":"akamai"}]}`)

	res, err := s.GetUsage(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(*seen) != 1 {
		t.Fatalf("got %d requests, want 1", len(*seen))
	}
	req := (*seen)[0]
	if req.Method != http.MethodGet || req.Header.Get("x-original-url") != usageURL {
		t.Errorf("sent %s %s, want GET %s", req.Method, req.Header.Get("x-original-url"), usageURL)
	}
	if req.Header.Get("x-api-key") != "test-key" {
		t.Errorf("x-api-key = %q", req.Header.Get("x-api-key"))
	}

	want := UsagePlan{
		Product:           "akamai",
		RequestsUsed:      44235934,
		RequestsQuota:     9999999999,
		RequestsRemaining: 9955764065,
		ExpiresAt:         time.Date(2027, 7, 31, 0, 0, 0, 0, time.UTC),
	}
	if res.BalanceEuro != 12.5 || len(res.Plans) != 1 || res.Plans[0] != want {
		t.Errorf("got %+v, want balance 12.5 and plan %+v", res, want)
	}
}

func TestGetUsageErrors(t *testing.T) {
	for _, tc := range []struct {
		status int
		body   string
		want   string
	}{
		{http.StatusForbidden, `{"error":"Invalid API key"}`, "Invalid API key"},
		{http.StatusInternalServerError, `{}`, "status code: 500"},
	} {
		s, _ := usageServer(t, tc.status, tc.body)
		_, err := s.GetUsage(context.Background())
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Errorf("status %d: err = %v, want it to contain %q", tc.status, err, tc.want)
		}
	}
}
