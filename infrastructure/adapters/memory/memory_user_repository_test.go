package adapters

import (
	"errors"
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"testing"
)

func TestNewMemoryUserRepository(t *testing.T) {
	type testCase struct {
		name        string
		id          string
		expectedErr error
	}

	user, err := aggregates.NewUser("biatechauth01", "UC1001", "User 1",
		"Sur 1", "user1@bia.com", "0729139504", "admin", "admin1", "987654321")
	if err != nil {
		t.Fatal(err)
	}

	id := user.GetUserID()
	repo := MemoryUserRepository{
		users: map[string]aggregates.UserAggregate{
			id: user,
		},
	}

	testCases := []testCase{
		{
			name:        "no user by id",
			id:          "UC1001_",
			expectedErr: aggregates.ErrUserNotFound,
		},
		{
			name:        "user by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.Get(testCase.id)
			if !errors.Is(testCase.expectedErr, err) {
				t.Errorf("%s: expected error: %s, got: %s", testCase.name, testCase.expectedErr, err)
			}
		})
	}
}
