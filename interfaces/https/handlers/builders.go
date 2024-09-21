package handlers

import (
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/cassandradb"
	"github.com/biangacila/biatechauth1/interfaces/https/controllers"
)

type ControllerRepository struct {
	repoUser    repositories.UserRepository
	repoLogin   repositories.LoginRepository
	repoGeneric repositories.GenericRepository[any]
}

type ControllerServices struct {
	User    *services.UserServiceImpl
	Login   *services.LoginServiceImpl
	Generic *services.GenericServiceImpl[any]
}

type ControllerHandlers struct {
	userController    controllers.UserController
	loginController   controllers.LoginController
	genericController controllers.GenericController[any]
}

func NewControllerRepositoryWithCassandra() ControllerRepository {
	return ControllerRepository{
		repoUser:    cassandradb.NewCassandraUserRepository(),
		repoLogin:   cassandradb.NewCassandraLoginRepository(),
		repoGeneric: cassandradb.NewCassandraGenericRepository[any](""),
	}
}

type Builders struct {
	repositories ControllerRepository
	services     ControllerServices
	handlers     ControllerHandlers
}

func NewBuilders() *Builders {
	repos := NewControllerRepositoryWithCassandra()
	return &Builders{
		repositories: repos,
	}
}

func (b *Builders) User(repo repositories.UserRepository) *Builders {
	b.repositories.repoUser = repo
	return b
}

func (b *Builders) Login(repo repositories.LoginRepository) *Builders {
	b.repositories.repoLogin = repo
	return b
}

func (b *Builders) Generic(repo repositories.GenericRepository[any]) *Builders {
	b.repositories.repoGeneric = repo
	return b
}

func (b *Builders) BuildService() ControllerServices {
	b.services.User = services.NewUserServiceImpl(b.repositories.repoUser)
	b.services.Login = services.NewLoginServiceImpl(b.repositories.repoLogin, b.repositories.repoUser)
	b.services.Generic = services.NewGenericServiceImpl(b.repositories.repoGeneric)

	return b.services
}
func (b *Builders) BuildRepository() (repos struct {
	repoUser    repositories.UserRepository
	repoLogin   repositories.LoginRepository
	repoGeneric repositories.GenericRepository[any]
}) {
	repos.repoUser = b.repositories.repoUser
	repos.repoLogin = b.repositories.repoLogin
	repos.repoGeneric = b.repositories.repoGeneric

	return repos
}
func (b *Builders) Build() ControllerHandlers {
	b.BuildService()
	h := b.handlers
	h.userController = controllers.NewUserController(b.services.User)
	h.loginController = controllers.NewLoginController(b.services.Login)
	h.genericController = controllers.NewGenericController(b.services.Generic)

	b.handlers = h
	return b.handlers
}
