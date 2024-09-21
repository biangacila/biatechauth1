package dtos

type LoginDto struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
type LoginCheckTokenDto struct {
	Token string `json:"token" validate:"required"`
}
