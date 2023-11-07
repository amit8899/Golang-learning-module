package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
func TestAbs(t *testing.T) {
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-1)
	}
}
