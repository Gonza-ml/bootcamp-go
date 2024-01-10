package main

import (
	"app/internal"
	"app/internal/storage"
	"time"
)

func main() {
	studentsMap := storage.NewStudentMap(nil, 0)

	students := []internal.Student{
		{
			Name:     "Juan",
			Apellido: "Granizo",
			DNI:      12345678,
			Fecha:    time.Date(1999, time.January, 23, 0, 0, 0, 0, time.Local),
		},
		{
			Name:     "Moises",
			Apellido: "Mareas",
			DNI:      11223344,
			Fecha:    time.Date(1998, time.February, 12, 0, 0, 0, 0, time.Local),
		},
		{
			Name:     "Pepe",
			Apellido: "Mujica",
			DNI:      11222333,
			Fecha:    time.Date(1979, time.July, 13, 0, 0, 0, 0, time.Local),
		},
		{
			Name:     "Pedro",
			Apellido: "Pascal",
			DNI:      33243566,
			Fecha:    time.Date(1975, time.November, 5, 0, 0, 0, 0, time.Local),
		},
	}

	for _, s := range students {
		studentsMap.AddStudent(&s)
	}

	studentsMap.ShowDetailStudents()
}
