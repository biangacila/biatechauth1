package valueobjects

import (
	"errors"
	"net/mail"
)

// ValidateEmail email format
func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateToken token (simple example, can be more complex)
func ValidateToken(token string) error {
	if len(token) < 10 {
		return errors.New("token is too short")
	}
	return nil
}
