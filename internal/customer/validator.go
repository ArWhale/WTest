package customer

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

func GenderValidation() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		return s == strings.ToLower("MALE") || s == strings.ToLower("FEMALE")
	}
}

func BirthDateValidation() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		t := time.Now()
		if ok {
			if t.Year()-date.Year() > 60 {
				return false
			}
			if t.Year()-date.Year() < 18 {
				return false
			}
		}
		return true
	}
}
