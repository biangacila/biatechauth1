package valueobjects

import "errors"

var (
	ErrFieldEmpty = errors.New("field is empty")
	ErrFieldShort = errors.New("field is too short")
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if value == "" {
		return Name{}, ErrFieldEmpty
	}
	if len(value) < 3 {
		return Name{}, ErrFieldShort
	}
	return Name{value: value}, nil
}
func ValidName(value string) error {
	if _, err := NewName(value); err != nil {
		return err
	}
	return nil
}
