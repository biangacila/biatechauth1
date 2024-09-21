package dtos

type UserPayloadDto struct {
	Password   string `json:"password" validate:"required,min=3,max=130"`
	GivenName  string `json:"given_name" validate:"required,min=3,max=130"`
	FamilyName string `json:"family_name" validate:"required,min=3,max=130"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone_number" validate:"required,number,min=9,max=11"`
}

type UserPayloadLockDto struct {
	Username string `json:"username" validate:"required,email"`
}
