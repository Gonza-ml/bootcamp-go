package internal

import "time"

type Student struct {
	IdStudent int
	Name      string
	Apellido  string
	DNI       int
	Fecha     time.Time
}
