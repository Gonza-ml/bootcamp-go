package salary

import (
	"errors"
	"fmt"
)

var Error1 = errors.New("Error: salary is less than 10000")

type ErrorSalary struct {
	msg string
}

func (e *ErrorSalary) Error() string {
	return fmt.Sprintf(e.msg)
}

func CalcularImpuesto(n float64) (result float64, err error) {
	switch {
	case n == 0:
		err = &ErrorSalary{msg: "Error salario 0"}
	case n < 10000:
		err = Error1
	case n < 50000:
		err = &ErrorSalary{msg: "Error: the salary entered does not reach the taxable minimum"}
	case n > 150000:
		result += n * 0.27
		err = &ErrorSalary{msg: "Must pay tax"}
	case n > 50000:
		result += n * 0.17
		err = &ErrorSalary{msg: "Must pay tax"}
	}
	return
}

func CalcularSalario(mn int, cat string) (salary float64, err string) {
	if mn <= 0 {
		return 0, "Tiempo de trabajo invalido"
	}
	switch cat {
	case "A":
		salary := float64(3000 * (mn / 60))
		salary += salary * 0.50
		return salary, ""
	case "B":
		salary := float64(1500 * (mn / 60))
		salary += salary * 0.20
		return salary, ""
	case "C":
		salary := float64(1000 * (mn / 60))
		return salary, ""
	default:
		return 0, "Categoria invalida"
	}
}
