package services

import (
	"github.com/biangacila/biatechauth1/application/dtos"
	"golang.org/x/oauth2"
)

type AuthServiceImpl struct {
}

func (a AuthServiceImpl) StoreTokenRealtime(token string, user map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) RegisterToken(sessionId string, token *oauth2.Token, provider string, user interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) CreateLogin(login dtos.LoginDto) error {
	//TODO implement me
	panic("implement me")
}
