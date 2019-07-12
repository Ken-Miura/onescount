package onescount

import (
	"math"
	"math/bits"
	"testing"
)

func TestOnesCount(t *testing.T) {
	var onesCountTests = []struct {
		n        uint64 // input
		expected int    // expected
	}{
		{0x1234567890ABCDEF, bits.OnesCount64(0x1234567890ABCDEF)},
		{0, bits.OnesCount64(0)},
		{1, bits.OnesCount64(1)},
		{2, bits.OnesCount64(2)},
		{3, bits.OnesCount64(3)},
		{1234567890987654321, bits.OnesCount64(1234567890987654321)},
		{math.MaxUint64, bits.OnesCount64(math.MaxUint64)},
	}

	for _, tt := range onesCountTests {
		actual := OnesCount64(tt.n)
		if actual != tt.expected {
			t.Fatalf("expected is %d, but actual is %d.", tt.expected, actual)
		}
	}
}

func BenchmarkOnesCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnesCount64(0x1234567890ABCDEF)
	}
}

func BenchmarkStdLibOnesCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.OnesCount64(0x1234567890ABCDEF)
	}
}
