package main

import (
	"app/internal/school"
	"fmt"
)

func main() {
	qualifications := []int{10, 6, 8, 7, 5}
	avg, err := school.Average(qualifications...)
	fmt.Println("El promedio de las notas:", qualifications, " es: ", avg)

	qualifications = []int{10, 10, 10, 10}
	avg, err = school.Average(qualifications...)
	fmt.Println("El promedio de las notas:", qualifications, " es: ", avg)
}
