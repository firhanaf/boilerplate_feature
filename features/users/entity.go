package users

import "time"

type UserCore struct {
	ID        string
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Login struct {
	Identifier string
	Password   string
}

type UserDataInterface interface {
	Login(Identifier string, Password string) (*UserCore, error)
	Register(Input UserCore) error
	UpdateProfile(ID string, Input UserCore) error
	GetProfile(ID string) (*UserCore, error)
	GetAllUsers() ([]UserCore, error)
	DeleteAccount(ID string) error
}

type UserServiceInterface interface {
	Login(Identifier string, Password string) (string, error)
	Register(Input UserCore) error
	UpdateProfile(ID string, Input UserCore) error
	GetProfile(ID string) (*UserCore, error)
	GetAllUsers() ([]UserCore, error)
	DeleteAccount(ID string) error
}
