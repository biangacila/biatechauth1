package entities

import (
	"fmt"
	"github.com/biangacila/biatechauth1/domain/valueobjects"
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

	Provider string // local, google, facebox
	Id       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) String() string {
	return valueobjects.NameToTitleCase(fmt.Sprintf("%v %v", u.GivenName, u.FamilyName))
}
