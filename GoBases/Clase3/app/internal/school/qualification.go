package school

import (
	"math"
)

type QualificationOperator string

type QualificationOperation func(...float64) (float64, string)

const (
	MinimumConst = "minimum"
	MaximumConst = "maximum"
	AverageConst = "average"
)

// Average return a average for qualifications
func Average(qa ...float64) (avg float64, err string) {
	var suma float64
	if len(qa) == 0 {
		err = "Notas vacias"
		return
	}
	for _, q := range qa {
		suma += q
	}
	avg = float64(suma) / float64(len(qa))
	return
}

// Minimum return a Minimum for qualifications
func Minimum(qa ...float64) (minimum float64, err string) {
	minimum = math.MaxFloat64
	if len(qa) == 0 {
		err = "Notas vacias"
		return
	}
	for _, q := range qa {
		if q < minimum {
			minimum = q
		}
	}
	return
}

// Maximum return a Maximum for qualifications
func Maximum(qa ...float64) (maximum float64, err string) {
	maximum = math.MinInt
	for _, q := range qa {
		if q > maximum {
			maximum = q
			return
		}
	}
	return
}
