package valueobjects

import (
	"fmt"
	"regexp"
)

// PhoneNumber represents a phone number value object
type PhoneNumber struct {
	number string
}

// NewPhoneNumber creates a new PhoneNumber object and validates the format
func ValuePhoneNumber(number string) (PhoneNumber, error) {
	// Define the regex patterns for valid phone number formats
	localFormat := `^0[0-9]{9}$`          // 10 digits, starting with 0
	internationalFormat := `^27[0-9]{9}$` // 11 digits, starting with country code 27

	// Compile the regex patterns
	localRegex := regexp.MustCompile(localFormat)
	internationalRegex := regexp.MustCompile(internationalFormat)

	// Validate the phone number format
	switch len(number) {
	case 10:
		// Local format (10 digits starting with 0)
		if !localRegex.MatchString(number) {
			return PhoneNumber{}, fmt.Errorf("invalid 10-digit phone number format: %s", number)
		}
	case 11:
		// International format (11 digits starting with 27 for South Africa)
		if !internationalRegex.MatchString(number) {
			return PhoneNumber{}, fmt.Errorf("invalid 11-digit phone number format: %s", number)
		}
	default:
		// Invalid length
		return PhoneNumber{}, fmt.Errorf("phone number must be either 10 or 11 digits: %s", number)
	}

	// If valid, return a new PhoneNumber object
	return PhoneNumber{number: number}, nil
}

// String returns the phone number as a string
func (p PhoneNumber) String() string {
	return p.number
}
