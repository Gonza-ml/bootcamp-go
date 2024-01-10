package main

import (
	"app/internal/salary"
	"fmt"
)

func main() {
	empleado := "Empleado A"
	categoria := "A"
	minutos := 1200

	salario, err := salary.CalcularSalario(minutos, categoria)
	if err != "" {
		fmt.Println("Error:", err)
	}
	fmt.Printf("El salario del %s es de %v \n", empleado, salario)

	empleado = "Empleado B"
	categoria = "B"
	minutos = 1200

	salario, err = salary.CalcularSalario(minutos, categoria)
	if err != "" {
		fmt.Println("Error:", err)
	}
	fmt.Printf("El salario del %s es de %v \n", empleado, salario)

	empleado = "Empleado C"
	categoria = "C"
	minutos = 1200

	salario, err = salary.CalcularSalario(minutos, categoria)
	if err != "" {
		fmt.Println("Error:", err)
	}
	fmt.Printf("El salario del %s es de %v \n", empleado, salario)
}
