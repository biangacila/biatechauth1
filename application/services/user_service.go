package services

import (
	"github.com/biangacila/biatechauth1/domain/entities"
)

type UserService interface {
	GenericService[entities.User]
	Create(name, surname, email, phone, password, id, provider, picture string, verifiedEmail bool) (entities.User, error)
	Lock(email string) error
	UnLock(email string) error
	UserExists(email string) (entities.User, error)
}
