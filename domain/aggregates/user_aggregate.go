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

func (c *UserAggregate) NewUser(email, name, surname, phone, password, id, provider, picture string, verifiedEmail bool) (user entities.User, err error) {
	newPassword := password
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
	if provider == "local" {
		if err = valueobjects.ValidPasswordPolicy(password); err != nil {
			return entities.User{}, err
		}
		password, err = valueobjects.NewPassword(password)
		if err != nil {
			return entities.User{}, err
		}
		newPassword, err = valueobjects.HashPassword(password)
		if err != nil {
			return entities.User{}, err
		}
	}

	user = entities.User{
		Email:         email,
		Phone:         phone,
		VerifiedEmail: verifiedEmail,
		GivenName:     name,
		FamilyName:    surname,
		Picture:       picture,
		Locale:        "",
		Password:      newPassword,
		Id:            id,
		Provider:      provider,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Status:        "active",
	}
	return user, nil
}
func (c *UserAggregate) HashPassword(password string) (newPassword string, err error) {
	if err = valueobjects.ValidPasswordPolicy(password); err != nil {
		return "", err
	}
	password, err = valueobjects.NewPassword(password)
	if err != nil {
		return "", err
	}
	return valueobjects.HashPassword(password)
}
