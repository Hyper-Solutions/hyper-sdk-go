package internal

import (
	"math/rand"
	"testing"
)

const correctOutput = "0.da49605db64e5"

func TestFloatToStringRadix16(t *testing.T) {
	buf := make([]byte, 18)
	length := FloatToStringRadix(0.852682135466517, 16, buf)
	v := string(buf[:length])
	if v != correctOutput {
		t.Errorf("expected %v, got: %v", correctOutput, v)
	}
}

func BenchmarkFloatToStringRadix16(b *testing.B) {
	buf := make([]byte, 18)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		FloatToStringRadix(rand.Float64(), 16, buf)
	}
}
