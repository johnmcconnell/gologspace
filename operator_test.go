package gologspace

import(
	"testing"
)

func BenchmarkOpAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := float64(i)
		a = a + a
	}
}

func BenchmarkOpSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := float64(i)
		a = a - a
	}
}

func BenchmarkOpMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := float64(i)
		a = a * a
	}
}

func BenchmarkOpDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := float64(i)
		a = a / a
	}
}
