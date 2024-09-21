package adapters

import (
	"fmt"
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"github.com/biangacila/biatechauth1/domain/valueobjects"
	"sync"
)

type MemoryUserRepository struct {
	users map[string]*aggregates.UserAggregate
	sync.Mutex
}

func (m *MemoryUserRepository) Save(aggregate *aggregates.UserAggregate) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) FindByID(id string) (*aggregates.UserAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) FindByEmail(id string) (*aggregates.UserAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) SaveLogin(userID string, login valueobjects.Login) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) GetLoginsByUserID(userID string) ([]valueobjects.Login, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryUserRepository) GetCurrentLogin(userID string) (*valueobjects.Login, error) {
	//TODO implement me
	panic("implement me")
}

func NewMemoryUserRepository() repositories.UserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*aggregates.UserAggregate),
	}
}

func (m *MemoryUserRepository) Add(aggregate aggregates.UserAggregate) error {
	m.Lock()
	defer m.Unlock()

	if m.users == nil {
		m.users = make(map[string]*aggregates.UserAggregate)
	}
	useId, err := aggregate.GetUserID()
	if err != nil {
		return err
	}
	if _, ok := m.users[useId]; ok {
		return aggregates.ErrUserAlreadyExists
	}

	m.users[useId] = &aggregate
	idU := m.users[useId]
	id, err := idU.GetUserID()
	if err != nil {
		return err
	}

	fmt.Println("AFTER ADDED> ", id)
	return nil
}

func (m *MemoryUserRepository) Get(id string) (aggregates.UserAggregate, error) {
	user, ok := m.users[id]
	if !ok {
		return aggregates.UserAggregate{}, aggregates.ErrUserNotFound
	}
	return *user, nil
}

func (m *MemoryUserRepository) Update(aggregate aggregates.UserAggregate) error {
	m.Lock()
	defer m.Unlock()
	if m.users == nil {
		m.users = make(map[string]*aggregates.UserAggregate)
	}
	useId, err := aggregate.GetUserID()
	if err != nil {
		return err
	}
	if _, ok := m.users[useId]; !ok {
		return aggregates.ErrUserNotFound
	}
	m.users[useId] = &aggregate
	return nil
}
