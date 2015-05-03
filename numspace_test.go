package gologspace

import(
	"testing"
)

func BenchmarkNumSpaceAdd(b *testing.B) {
	s := NumSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := float64(i)
		s.Add(a, a)
	}
}

func BenchmarkNumSpaceSub(b *testing.B) {
	s := NumSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := float64(i)
		s.Sub(a, a)
	}
}

func BenchmarkNumSpaceMul(b *testing.B) {
	s := NumSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := float64(i)
		s.Mul(a, a)
	}
}

func BenchmarkNumSpaceDiv(b *testing.B) {
	s := NumSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := float64(i)
		s.Div(a, a)
	}
}
