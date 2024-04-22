package akamai

import (
	"io"
	"strings"
	"testing"
)

func TestParseScriptPath(t *testing.T) {
	path, err := ParseScriptPath(strings.NewReader(`<script type="text/javascript"  src="/QaZkxg1hkkBeBuc7J_KFVltz/tEz56cktSmc7/XEcFAQ/Aj4OUi/EqXkc"></script>`))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(path)
}

func BenchmarkParseScriptPath(b *testing.B) {
	script := strings.NewReader(`<script type="text/javascript"  src="/QaZkxg1hkkBeBuc7J_KFVltz/tEz56cktSmc7/XEcFAQ/Aj4OUi/EqXkc"></script>`)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := ParseScriptPath(script)
		if err != nil {
			b.Fatal(err)
		}

		script.Seek(0, io.SeekStart)
	}
}
