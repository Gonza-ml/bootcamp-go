package main

import (
	"app/internal/school"
	"fmt"
)

func main() {
	minFunc, err := school.Orchestrator(school.MinimumConst)
	if err != "" {
		fmt.Println(err)
	}
	maxFunc, err := school.Orchestrator(school.MaximumConst)
	if err != "" {
		fmt.Println(err)
	}
	avgFunc, err := school.Orchestrator(school.AverageConst)
	if err != "" {
		fmt.Println(err)
	}

	notas := make([]float64, 0)
	notas = append(notas, 1, 2, 3, 4, 5, 10, 10, 10, 10, 10)

	valorMin, err := minFunc(notas...)
	if err != "" {
		fmt.Println(err)
	}
	valorMax, err := maxFunc(notas...)
	if err != "" {
		fmt.Println(err)
	}
	valorAvg, err := avgFunc(notas...)
	if err != "" {
		fmt.Println(err)
	}

	fmt.Println("La nota Minima es", valorMin)
	fmt.Println("La nota Maxima es", valorMax)
	fmt.Println("La nota Promedio es", valorAvg)
}
