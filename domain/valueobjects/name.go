package valueobjects

import (
	"errors"
	"strings"
	"unicode"
)

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

// Alternative using a custom function to ensure correct behavior
func NameToTitleCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
