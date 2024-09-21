package valueobjects

import (
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	value string
}
type ID string

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func NewEmail(value string) (Email, error) {
	if !emailRegex.MatchString(value) {
		return Email{}, fmt.Errorf("invalid email format")
	}
	value = strings.ToLower(value)
	value = strings.TrimSpace(value)
	return Email{value: value}, nil
}
func FormatEmail(email string) string {
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)
	return email
}
func (e Email) String() string {
	return e.value
}
