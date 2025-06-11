package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Role      string
}
