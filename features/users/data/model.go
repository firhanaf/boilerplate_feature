package data

import (
	"boilerplate-feature/features/users"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique"`
	Password  string
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Role      string
}

func UserModeltoCore(input User) users.UserCore {
	return users.UserCore{
		ID:        input.ID,
		Username:  input.Username,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Role:      input.Role,
	}
}

func UserCoretoModel(input users.UserCore) User {
	return User{
		Username:  input.Username,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Role:      input.Role,
	}
}
