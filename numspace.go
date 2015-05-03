package gologspace

type NumSpace struct{}

func (l *NumSpace) Enter(x float64) float64 {
	return x
}

func (l *NumSpace) Exit(x float64) float64 {
	return x
}

func (l *NumSpace) Add(x, y float64) float64 {
	return x + y
}

func (l *NumSpace) Div(x, y float64) float64 {
	return x / y
}

func (l *NumSpace) Mul(x, y float64) float64 {
	return x * y
}

func (l *NumSpace) Sub(x, y float64) float64 {
	return x - y
}
