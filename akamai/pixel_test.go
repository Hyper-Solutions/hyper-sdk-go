package akamai

import (
	"io"
	"strings"
	"testing"
)

func TestParsePixelHtmlVar(t *testing.T) {
	if v, err := ParsePixelHtmlVar(strings.NewReader(`<script >bazadebezolkohpepadr="210295263"</script>`)); err != nil {
		t.Fatal(err)
	} else {
		t.Log(v)
	}
}

func TestParsePixelScriptURL(t *testing.T) {
	if scriptUrl, postUrl, err := ParsePixelScriptURL(strings.NewReader(`src="https://www.example.com/akam/13/c88db65"`)); err != nil {
		t.Fatal(err)
	} else {
		t.Log(scriptUrl, postUrl)
	}
}

func BenchmarkParsePixelScriptURL(b *testing.B) {
	script := strings.NewReader(`src="https://www.example.com/akam/13/c88db65"`)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := ParsePixelScriptURL(script)
		if err != nil {
			b.Fatal(err)
		}

		script.Seek(0, io.SeekStart)
	}
}
