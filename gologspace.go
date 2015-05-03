package gologspace

import(
	"math"
)

func LogAdd(logX, logY float64) float64 {
	if logY > logX {
		temp := logX
		logX = logY
		logY = temp
	}
	if logX == math.Inf(-1) {
		return logX
	}
	negDiff := logY - logX
	if negDiff < 20.0 {
		return logX
	}
	return logX + math.Log(1.0 + math.Exp(negDiff))
}
