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

func TestParseSbsdScriptPath(t *testing.T) {
	path, err := ParseSbsdScriptPath(strings.NewReader(`<script type="text/javascript"  src="/QaZkxg1hkkBeBuc7J_KFVltz/tEz56cktSmc7/XEcFAQ/Aj4OUi/EqXkc?v=83cfdce1-9f29-c905-41c3-cf81f625170f"></script>`))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(path)
}

func BenchmarkParseSbsdScriptPath(b *testing.B) {
	script := strings.NewReader(`<script type="text/javascript"  src="/QaZkxg1hkkBeBuc7J_KFVltz/tEz56cktSmc7/XEcFAQ/Aj4OUi/EqXkc?v=83cfdce1-9f29-c905-41c3-cf81f625170f"></script>`)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := ParseSbsdScriptPath(script)
		if err != nil {
			b.Fatal(err)
		}

		script.Seek(0, io.SeekStart)
	}
}

