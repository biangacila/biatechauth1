package services

import (
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"github.com/biangacila/biatechauth1/internal/utils"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	GenericServiceImpl[entities.User]
	repo repositories.UserRepository
}

// NewUserServiceImpl creates a new instance of UserServiceImpl
func NewUserServiceImpl(repo repositories.UserRepository) *UserServiceImpl {
	var genericService = NewGenericServiceImpl(repo)
	return &UserServiceImpl{
		GenericServiceImpl: *genericService, // Use a pointer receiver for GenericServiceImpl
		repo:               repo,
	}
}

// Create creates a new user
func (u *UserServiceImpl) Create(name, surname, email, phone, password, id, provider, picture string, verifiedEmail bool) (entities.User, error) {
	agg := aggregates.NewUserAggregate()
	user, err := agg.NewUser(email, name, surname, phone, password, id, provider, picture, verifiedEmail)
	if err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		utils.NewLoggerSlog().Debug("Create user error", "problem", err.Error())
		return entities.User{}, err
	}
	// TODO please create login of this user to auth micro service

	return user, u.Save("users", user, entities.User{})
}

// UserExists checks if a user exists by their user ID (or any other criteria)
func (u *UserServiceImpl) UserExists(email string) (entities.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Lock locks a user account
func (u *UserServiceImpl) Lock(email string) (err error) {
	if _, err = u.repo.FindByEmail(email); err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return err
	}
	if err = u.repo.Lock(email); err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return err
	}
	return nil
}

// UnLock unlocks a user account
func (u *UserServiceImpl) UnLock(email string) (err error) {
	if _, err = u.repo.FindByEmail(email); err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return err
	}
	if err = u.repo.UnLock(email); err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return err
	}
	return nil
}
