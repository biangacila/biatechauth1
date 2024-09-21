package valueobjects

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
func ValidPasswordPolicy(password string) error {
	// Set your password policy rules here
	minLength := 8
	maxLength := 20

	// Validate password length
	if len(password) < minLength {
		return fmt.Errorf("password must be at least %d characters long", minLength)
	}
	if len(password) > maxLength {
		return fmt.Errorf("password must be no more than %d characters long", maxLength)
	}

	// Initialize flags for password rules
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecialChar := false

	// Iterate over each character in the password to check policy rules
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	// Validate required character types
	if !hasUppercase {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLowercase {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return fmt.Errorf("password must contain at least one digit")
	}
	if !hasSpecialChar {
		return fmt.Errorf("password must contain at least one special character")
	}

	// If the password passes all validations, return the PasswordPolicy object
	return nil
}
func NewPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	if strings.Contains(password, "'") {
		return "", errors.New("password must not contain ''")
	}
	return password, nil
}
