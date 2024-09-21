package services

import (
	"github.com/biangacila/biatechauth1/domain/entities"
	"time"
)

type LoginService interface {
	NewLogin(username, password string) (entities.User, string, error)
	HasLogin(username string) (time.Time, bool, error)
	IsValueToken(token string) (time.Time, bool, error)
}
