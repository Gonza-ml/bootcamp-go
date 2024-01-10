package internal

// StorageStudent is an interface for storing students
type StorageStudent interface {
	AddStudent(student *Student) (err string)
	ShowDetailStudents()
}
