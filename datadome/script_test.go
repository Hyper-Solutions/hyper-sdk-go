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
			name: "captcha deferred with build prefix",
			html: `<script defer src="https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js"></script>`,
			want: "https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js",
		},
		{
			name: "build prefix after the inline dd object",
			html: `<script data-cfasync="false">var dd={'rt':'c'}</script><script data-cfasync="false" src="https://ct.captcha-delivery.com/c.js"></script><script defer src="https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js"></script>`,
			want: "https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js",
		},
		{
			name: "any file name",
			html: `<script defer src="https://ct.captcha-delivery.com/interstitial-next.js"></script>`,
			want: "https://ct.captcha-delivery.com/interstitial-next.js",
		},
		{
			name: "defer after src",
			html: `<script src="https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js" defer></script>`,
			want: "https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js",
		},
		{
			name: "first deferred tag wins",
			html: `<script defer src="https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js"></script><script defer src="https://example.com/other.js"></script>`,
			want: "https://ct.captcha-delivery.com/captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js",
		},
		{
			name: "defer in src is not the attribute",
			html: `<script src="https://example.com/defer.js"></script>`,
			want: "",
		},
		{
			name: "loader script without defer",
			html: `<script src="https://ct.captcha-delivery.com/i.js"></script>`,
			want: "",
		},
		{
			name: "single quotes",
			html: `<script defer src='https://ct.captcha-delivery.com/captcha.1.34.0.202609141.js'></script>`,
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
