package user

import "gorm.io/gorm"

var DB *gorm.DB
var error error

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
