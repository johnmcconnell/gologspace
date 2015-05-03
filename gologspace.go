package gologspace

import(
	"math"
)

func LogAddAll(logX float64, logYs ...float64) float64 {
	result := logX
	for _, logY := range logYs {
		result = LogAdd(result, logY)
	}
	return result
}

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
	if negDiff < -20.0 {
		return logX
	}
	return logX + math.Log(1.0 + math.Exp(negDiff))
}

func LogDivAll(logX float64, logYs ...float64) float64 {
	result := logX
	for _, logY := range logYs {
		result = LogDiv(result, logY)
	}
	return result
}

func LogDiv(logX, logY float64) float64 {
	return logX - logY
}

func LogMultAll(logX float64, logYs ...float64) float64 {
	result := logX
	for _, logY := range logYs {
		result = LogMult(result, logY)
	}
	return result
}

func LogMult(logX, logY float64) float64 {
	return logX + logY
}

func LogProb(p float64) float64 {
	return math.Log(p)
}
