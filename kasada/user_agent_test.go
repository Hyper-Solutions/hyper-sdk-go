package kasada

import "testing"

func TestGenerateUserAgent(t *testing.T) {
	got := GenerateUserAgent("18.0.1", "com.example.app", "9.15.0")
	want := "7016adc4ba8d433533954ebad33be424"
	if got != want {
		t.Fatalf("GenerateUserAgent = %q, want %q", got, want)
	}
}
