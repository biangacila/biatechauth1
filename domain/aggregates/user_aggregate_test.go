package aggregates

import (
	"errors"
	"testing"
)

func TestNewUser(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: ErrNameEmpty,
		},
		{
			test:        "Valid name",
			name:        "Merveilleux",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewUser("biatechauth01", "UC1001", tc.name, "Biangacila",
				"biangacila@gmail.com", "27729139504", "super", "aDmin2024*",
				"7611026059185")
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
