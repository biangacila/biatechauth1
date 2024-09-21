package entities

import (
	"fmt"
	"time"
)

type User struct {
	Email         string
	Phone         string
	VerifiedEmail bool
	GivenName     string
	FamilyName    string
	Picture       string
	Locale        string
	Password      string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u User) String() string {
	return fmt.Sprintf("%v %v", u.GivenName, u.FamilyName)
}
