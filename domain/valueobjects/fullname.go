package valueobjects

import "fmt"

type FullName struct {
	Name    string
	Surname string
}

func NewFullName(name string, surname string) *FullName {
	return &FullName{Name: name, Surname: surname}
}

func (f *FullName) ToString() string {
	return fmt.Sprintf("%v %v", f.Name, f.Surname)
}
