package utils

import (
	"time"
)

func GetHours(fechaInicio time.Time, fechaFinal time.Time) bool {
	duration := fechaInicio.Sub(fechaFinal)
	elapsed := duration.Hours() / 24
	return elapsed > 31 || false
}
