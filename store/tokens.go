package store

import (
	"errors"
	"sync"
	"time"
)

type storeToken struct {
	Username  string
	Token     string
	ExpiredAt time.Time
}

type ServeStore struct {
	sync.RWMutex
	Tokens map[string]storeToken
}

func NewServeStore() *ServeStore {
	return &ServeStore{Tokens: make(map[string]storeToken)}
}

var stores *ServeStore

func InitTokens() {
	stores = NewServeStore()
}
func CloseTokens() {
	if stores != nil {
		stores.Tokens = make(map[string]storeToken)
	}
}
func (c *ServeStore) AddToken(username, token string, expiredAt time.Time) error {
	c.Lock()
	defer c.Unlock()
	c.Tokens[token] = storeToken{
		Username:  username,
		Token:     token,
		ExpiredAt: expiredAt,
	}
	return nil
}
func (c *ServeStore) RemoveToken(token string) error {
	c.Lock()
	defer c.Unlock()
	delete(c.Tokens, token)
	return nil
}
func (c *ServeStore) FindToken(token string) (storeToken, error) {
	c.RLock()
	defer c.RUnlock()
	t, ok := c.Tokens[token]
	if !ok {
		return t, errors.New("token not found")
	}
	return t, nil
}
func (c *ServeStore) IsValidToken(token string) error {
	c.RLock()
	defer c.RUnlock()
	t, err := c.FindToken(token)
	if err != nil {
		return err
	}
	if t.ExpiredAt.After(time.Now()) {
		// let remove it from our store
		delete(c.Tokens, token)
		return errors.New("token is expired")
	}
	return nil
}

func NewToken(username, token string, expiredAt time.Time) error {
	return stores.AddToken(username, token, expiredAt)
}
func GetStore() *ServeStore {
	if stores == nil {
		InitTokens()
	}
	return stores
}
