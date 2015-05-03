package gologspace

import(
	"math"
)

type LogSpace struct{}

func (l LogSpace) Enter(x float64) float64 {
	return math.Log(x)
}

func (l LogSpace) Exit(x float64) float64 {
	return math.Exp(x)
}

func (l LogSpace) Add(logX, logY float64) float64 {
	if logY > logX {
		temp := logX
		logX = logY
		logY = temp
	}
	if logX == math.Inf(-1) {
		return logX
	}
	negDiff := logY - logX
	if negDiff < -20.0 {
		return logX
	}
	return logX + math.Log(1.0 + math.Exp(negDiff))
}

func (l LogSpace) Div(logX, logY float64) float64 {
	return logX - logY
}

func (l LogSpace) Mul(logX, logY float64) float64 {
	return logX + logY
}

func (l LogSpace) Sub(logX, logY float64) float64 {
	negDiff := logY - logX
	return logX + math.Log(-1.0 + math.Exp(negDiff))
}
