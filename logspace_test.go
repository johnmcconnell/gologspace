
package gologspace

import(
	"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

const(
	Angstrom = 1E-10
)

func BenchmarkLogSpaceAdd(b *testing.B) {
	s := LogSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := s.Enter(float64(i))
		s.Add(a, a)
	}
}

func BenchmarkLogSpaceSub(b *testing.B) {
	s := LogSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := s.Enter(float64(i))
		s.Sub(a, a)
	}
}

func BenchmarkLogSpaceMul(b *testing.B) {
	s := LogSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := s.Enter(float64(i))
		s.Mul(a, a)
	}
}

func BenchmarkLogSpaceDiv(b *testing.B) {
	s := LogSpace{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := s.Enter(float64(i))
		s.Div(a, a)
	}
}

func TestAdd(t *testing.T) {
	s := LogSpace{}
	L := 100
	for x := 0; x < L; x += 10 {
		xA := float64(2 * x + 1)
		yA := float64(x - 25)
		A := xA + yA
		xB := s.Enter(xA)
		yB := s.Enter(yA)
		B := s.Exit(s.Add(xB, yB))
		assert.InDelta(t, A, B, Angstrom, fmt.Sprintf("Add(%v, %v)", xB, yB))
	}
}

func TestSub(t *testing.T) {
	s := LogSpace{}
	L := 100
	for x := 0; x < L; x += 10 {
		xA := float64(2 * x)
		yA := float64(x - 25)
		A := xA - yA
		xB := s.Enter(xA)
		yB := s.Enter(yA)
		B := s.Exit(s.Sub(xB, yB))
		assert.InDelta(t, A, B, Angstrom, fmt.Sprintf("Sub(%v,%v)", xA, yA))
	}
}

func TestMul(t *testing.T) {
	s := LogSpace{}
	L := 100
	for x := 0; x < L; x += 10 {
		xA := float64(2 * x)
		yA := float64(x - 25)
		A := xA * yA
		xB := s.Enter(xA)
		yB := s.Enter(yA)
		B := s.Exit(s.Mul(xB, yB))
		assert.InDelta(t, A, B, Angstrom, fmt.Sprintf("Mul(%v,%v)", xA, yA))
	}
}

func TestDiv(t *testing.T) {
	s := LogSpace{}
	L := 100
	for x := 0; x < L; x += 10 {
		xA := float64(2 * x)
		yA := float64(x - 25)
		A := xA / yA
		xB := s.Enter(xA)
		yB := s.Enter(yA)
		B := s.Exit(s.Div(xB, yB))
		assert.InDelta(t, A, B, Angstrom, fmt.Sprintf("Div(%v,%v)", xA, yA))
	}
}
