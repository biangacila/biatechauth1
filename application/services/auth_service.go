package services

import "github.com/biangacila/biatechauth1/application/dtos"

type AuthService interface {
	RegisterToken(token string, provider string, user interface{}) error
	CreateLogin(login dtos.LoginDto) error
	StoreTokenRealtime(token string, user map[string]interface{}) error
}
