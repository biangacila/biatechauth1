package aggregates

import (
	"errors"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/biangacila/biatechauth1/domain/valueobjects"
	"time"
)

var (
	ErrNameEmpty    = errors.New("aggregate: name is empty")
	ErrSurnameEmpty = errors.New("aggregate: surname is empty")
	ErrEmailEmpty   = errors.New("aggregate: email is empty or invalid")
)

type UserAggregate struct {
	user *entities.User
}

func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

func (c *UserAggregate) NewUser(email, name, surname, phone, password string) (user entities.User, err error) {
	email = valueobjects.FormatEmail(email)
	if _, err = valueobjects.NewName(name); err != nil {
		return entities.User{}, ErrNameEmpty
	}
	if _, err = valueobjects.NewName(surname); err != nil {
		return entities.User{}, ErrSurnameEmpty
	}
	if _, err = valueobjects.NewEmail(email); err != nil {
		return entities.User{}, ErrEmailEmpty
	}
	if _, err = valueobjects.ValuePhoneNumber(phone); err != nil {
		return entities.User{}, err
	}
	if err = valueobjects.ValidPasswordPolicy(password); err != nil {
		return entities.User{}, err
	}

	password, err = valueobjects.NewPassword(password)
	if err != nil {
		return entities.User{}, err
	}

	newPassword, err := valueobjects.HashPassword(password)
	if err != nil {
		return entities.User{}, err
	}

	user = entities.User{
		Email:         email,
		Phone:         phone,
		VerifiedEmail: false,
		GivenName:     name,
		FamilyName:    surname,
		Picture:       "",
		Locale:        "",
		Password:      newPassword,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Status:        "active",
	}
	return user, nil
}
