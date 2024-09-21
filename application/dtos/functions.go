package dtos

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

func Validate[T any](data any, t T) error {
	validate := validator.New()
	return validate.Struct(data)
}
func ToEntity[T any](entity any, t T) T {
	str, _ := json.Marshal(entity)
	var out T
	_ = json.Unmarshal(str, &out)
	return out
}
func ToEntities[T any](entity any, t []T) []T {
	str, _ := json.Marshal(entity)
	var out []T
	_ = json.Unmarshal(str, &out)
	return out
}
func ValidateAnyWithAnyDto[T any](data any, t T) error {
	newRec := ToEntity(data, t)
	validate := validator.New()
	return validate.Struct(newRec)
}
