package aggregates

import (
	"errors"
	"fmt"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/biangacila/biatechauth1/domain/valueobjects"
	"time"
)

type LoginAggregate struct {
	login *entities.Login
}

func NewLoginAggregate() *LoginAggregate {
	return &LoginAggregate{}
}

func (c *LoginAggregate) New(username, token string, expiredAt time.Time) (*entities.Login, error) {
	if err := valueobjects.ValidateEmail(username); err != nil {
		return nil, err
	}
	if err := valueobjects.ValidateToken(token); err != nil {
		return nil, err
	}
	login := &entities.Login{
		Username:    username,
		SignedToken: token,
		Updated:     time.Now(),
		ExpiredAt:   expiredAt,
	}
	c.login = login
	return login, nil
}
func (c *LoginAggregate) Set(login *entities.Login) {
	c.login = login
}

func (c *LoginAggregate) ValidUsernamePassword(username, password string) (err error) {
	if err = valueobjects.ValidateEmail(username); err != nil {
		return err
	}
	if err = valueobjects.ValidPasswordPolicy(password); err != nil {
		return err
	}
	return nil
}

// CreateToken when login success
func (c *LoginAggregate) CreateToken(email, givenName, familyName, phone string) (signedToken, signedFreshToken string, err error) {
	if err = valueobjects.ValidateEmail(email); err != nil {
		return "", "", err
	}
	if _, err = valueobjects.ValuePhoneNumber(phone); err != nil {
		return "", "", err
	}
	if err = valueobjects.ValidName(givenName); err != nil {
		return "", "", fmt.Errorf("aggregate given name: %v", err.Error())
	}
	if err = valueobjects.ValidName(familyName); err != nil {
		return "", "", fmt.Errorf("aggregate family name: %v", err.Error())
	}
	return valueobjects.GenerateAllTokens(email, givenName, familyName, phone)
}

// UpdateToken token and expiration time
func (c *LoginAggregate) UpdateToken(newToken string, newExpiredAt time.Time) error {
	if err := valueobjects.ValidateToken(newToken); err != nil {
		return err
	}
	l := c.login
	l.SignedToken = newToken
	l.ExpiredAt = newExpiredAt
	l.Updated = time.Now()
	c.login = l
	return nil
}
func (c *LoginAggregate) Matches(hashed, plainer string) (err error) {
	if ok := valueobjects.ComparePasswords(hashed, []byte(plainer)); !ok {
		return errors.New("invalid login credentials, please check your username or password")
	}
	return nil
}
func (c *LoginAggregate) GetTokenExpiryAndValidity(token string) (expiredAt time.Time, bool bool, err error) {
	return valueobjects.GetTokenExpiryAndValidity(token)
}
