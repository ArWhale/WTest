package customer

import (
	"time"
)

type Customer struct {
	ID        int64     `form:"id" json:"id,omitempty" binding:"-"`
	FirstName string    `form:"firstName" json:"firstName" binding:"required"`
	LastName  string    `form:"lastName" json:"lastName" binding:"required"`
	Gender    string    `form:"gender" json:"gender" binding:"required,genderCustom"`
	Email     string    `form:"email" json:"email" binding:"required"`
	Address   string    `form:"address" json:"address" binding:"-"`
	BirthDate time.Time `form:"birthdate" json:"birthdate" binding:"required,birthdateCustom"`
}
