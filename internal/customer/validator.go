package customer

import (
	"github.com/ArWhale/WTest/internal/consts"
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
		s := fl.Field().String()
		dTime, err := time.Parse(consts.DefaultDateLayout, s)
		if err != nil {
			return false
		}
		t := time.Now()

		if t.Year()-dTime.Year() > 60 {
			return false
		}
		if t.Year()-dTime.Year() < 18 {
			return false
		}

		return true
	}
}
