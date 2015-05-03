package gologspace

type Space interface {
	Enter(float64) float64
	Add(x,y float64) float64
	Sub(x, y float64) float64
	Mul(x,y float64) float64
	Div(x, y float64) float64
	Exit(float64) float64
}
