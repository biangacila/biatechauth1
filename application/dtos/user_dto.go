package dtos

import (
	"github.com/biangacila/biatechauth1/domain/entities"
	"time"
)

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

type UserResponseDto struct {
	GivenName  string `json:"given_name" `
	FamilyName string `json:"family_name"`
	Email      string `json:"email" `
	Phone      string `json:"phone_number" `
	Picture    string `json:"picture" `
	CreatedAt  string `json:"created_at" `
	UpdatedAt  string `json:"updated_at"`
	Locale     string `json:"locale"`
	Status     string `json:"status" `
	Name       string `json:"name" `
}

// ToUserResponseDto Method to convert User to UserResponseDto
func ToUserResponseDto(u entities.User) UserResponseDto {
	return UserResponseDto{
		GivenName:  u.GivenName,
		FamilyName: u.FamilyName,
		Email:      u.Email,
		Phone:      u.Phone,
		Picture:    u.Picture,
		CreatedAt:  u.CreatedAt.Format(time.RFC3339), // Convert time.Time to string
		UpdatedAt:  u.UpdatedAt.Format(time.RFC3339), // Convert time.Time to string
		Locale:     u.Locale,
		Status:     u.Status,
		Name:       u.String(),
	}
}
