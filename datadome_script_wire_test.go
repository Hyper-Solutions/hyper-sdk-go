package hyper

import (
	"strings"
	"testing"

	"github.com/mailru/easyjson"
)

// models_easyjson.go is generated but committed, so a field missing from either
// half of the codec is dropped silently and looks like a server-side bug.
func TestDataDomeScriptFieldSurvivesTheWire(t *testing.T) {
	const script = "/* DataDome is a cyberfraud solution v1.33.0 */\n!function(){}();"

	t.Run("interstitial", func(t *testing.T) {
		raw, err := easyjson.Marshal(DataDomeInterstitialInput{Html: "<html></html>", Script: script})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(raw), `"script":`) {
			t.Fatalf("script missing from the encoded body: %s", raw)
		}
		var out DataDomeInterstitialInput
		if err := easyjson.Unmarshal(raw, &out); err != nil {
			t.Fatal(err)
		}
		if out.Script != script {
			t.Errorf("script did not round-trip:\n want %q\n got  %q", script, out.Script)
		}
	})

	t.Run("slider", func(t *testing.T) {
		raw, err := easyjson.Marshal(DataDomeSliderInput{Html: "<html></html>", Script: script})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(raw), `"script":`) {
			t.Fatalf("script missing from the encoded body: %s", raw)
		}
		var out DataDomeSliderInput
		if err := easyjson.Unmarshal(raw, &out); err != nil {
			t.Fatal(err)
		}
		if out.Script != script {
			t.Errorf("script did not round-trip:\n want %q\n got  %q", script, out.Script)
		}
	})

	t.Run("absent stays empty", func(t *testing.T) {
		var out DataDomeInterstitialInput
		if err := easyjson.Unmarshal([]byte(`{"userAgent":"UA","html":"<h/>","deviceLink":"https://x/","acceptLanguage":"en","ip":"1.2.3.4"}`), &out); err != nil {
			t.Fatal(err)
		}
		if out.Script != "" {
			t.Errorf("script = %q, want empty", out.Script)
		}
	})
}
