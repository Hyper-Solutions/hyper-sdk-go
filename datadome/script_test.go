package datadome

import "testing"

func TestParseChallengeScriptURL(t *testing.T) {
	cases := []struct {
		name string
		html string
		want string
	}{
		{
			name: "interstitial deferred",
			html: `<head><script defer src="https://ct.captcha-delivery.com/interstitial.1.33.0.202609141.js"></script></head>`,
			want: "https://ct.captcha-delivery.com/interstitial.1.33.0.202609141.js",
		},
		{
			name: "captcha deferred",
			html: `<script defer src="https://ct.captcha-delivery.com/captcha.1.34.0.202609141.js"></script>`,
			want: "https://ct.captcha-delivery.com/captcha.1.34.0.202609141.js",
		},
		{
			name: "interstitial deferred without build",
			html: `<head><script defer src="https://ct.captcha-delivery.com/interstitial.1.36.0.js"></script></head>`,
			want: "https://ct.captcha-delivery.com/interstitial.1.36.0.js",
		},
		{
			name: "captcha deferred without build",
			html: `<script defer src="https://ct.captcha-delivery.com/captcha.1.37.0.js"></script>`,
			want: "https://ct.captcha-delivery.com/captcha.1.37.0.js",
		},
		{
			name: "loader script without version",
			html: `<script src="https://ct.captcha-delivery.com/i.js"></script>`,
			want: "",
		},
		{
			name: "single quotes",
			html: `<script src='https://ct.captcha-delivery.com/captcha.1.34.0.202609141.js'></script>`,
			want: "https://ct.captcha-delivery.com/captcha.1.34.0.202609141.js",
		},
		{
			name: "inlined bundle has no tag",
			html: `<script>var ddm = {};/* DataDome is a cyberfraud solution v1.33.0 */!function(){}()</script>`,
			want: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := ParseChallengeScriptURL(c.html)
			if c.want == "" {
				if ok {
					t.Errorf("expected no script tag, got %q", got)
				}
				return
			}
			if !ok {
				t.Fatalf("expected to find %q", c.want)
			}
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
