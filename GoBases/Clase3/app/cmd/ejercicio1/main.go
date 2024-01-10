package main

import (
	"app/internal/salary"
	"errors"
	"fmt"
)

func main() {
	salario := 1000.0
	impuesto, err := salary.CalcularImpuesto(salario)
	if errors.Is(err, salary.Error1) {
		fmt.Println(err)
	}
	fmt.Printf("El impuesto al salario de %.0F es de %.0F \n", salario, impuesto)
}
