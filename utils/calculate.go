package utils

import (
	"fmt"
	"time"
)

func GetHours(fechaInicio time.Time, fechaFinal time.Time) bool {
	duration := fechaInicio.Sub(fechaFinal)
	fmt.Println(duration)
	elapsed := duration.Hours() / 24
	fmt.Println(elapsed)
	return elapsed > 31 || false

}
