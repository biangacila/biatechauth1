package repositories

import "github.com/biangacila/biatechauth1/domain/entities"

type LoginRepository interface {
	GenericRepository[entities.Login]
	New(login *entities.Login) (*entities.Login, error)
	HasLogin(username string) (*entities.Login, error)
}
