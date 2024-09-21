package repositories

import (
	"github.com/biangacila/biatechauth1/domain/entities"
)

type UserRepository interface {
	GenericRepository[entities.User]
	Lock(code string) error
	UnLock(email string) error
	FindByEmail(email string) (user entities.User, err error)
}
