package storage

import (
	"app/internal"
	"fmt"
)

// StudentMap is a storage for students that implements StorageStudent
type StudentMap struct {
	// students is a map of students
	// - key: student ID
	// - value: student
	students map[int]internal.Student

	// LastId is the last student id
	lastId int
}

func NewStudentMap(students map[int]internal.Student, lastId int) *StudentMap {
	defaultStudents := make(map[int]internal.Student)
	defaultLastId := 0
	if students != nil {
		defaultStudents = students
	}
	if lastId != 0 {
		defaultLastId = lastId
	}
	return &StudentMap{
		students: defaultStudents,
		lastId:   defaultLastId,
	}
}

// Ingresamos el parametro porm puntero porque vamos a actualizar el idStudent

func (s *StudentMap) AddStudent(student *internal.Student) (err string) {

	//check student
	if student.Name == "" {
		err = "Nombre invalido"
		return
	}

	if student.Apellido == "" {
		err = "Apellido invalido"
		return
	}

	if student.DNI == 0 {
		err = "DNI invalido"
		return
	}

	//check dni no repeat

	for _, st := range s.students {
		if student.DNI == st.DNI {
			err = "DNI repetido"
			return
		}
	}

	// id increment
	// (*s).lastId++
	s.lastId++

	//set id
	student.IdStudent = s.lastId

	// save student
	s.students[student.IdStudent] = *student
	return
}

func (s *StudentMap) ShowDetailStudents() {
	for _, student := range s.students {
		fmt.Println("Name: ", student.Name)
		fmt.Println("Apellido: ", student.Apellido)
		fmt.Println("DNI: ", student.DNI)
		fmt.Println("Fecha: ", student.Fecha.Format("02-01-2006"))
		fmt.Println()
	}
}
