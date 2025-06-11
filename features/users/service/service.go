package service

import (
	"boilerplate-feature/features/users"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func (u UserService) Login(Identifier string, Password string) (*users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Register(Input users.UserCore) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) UpdateProfile(Input users.UserCore) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetProfile(ID string) (*users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetAllUsers() ([]users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteAccount(ID string) error {
	//TODO implement me
	panic("implement me")
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
		validate: validator.New(),
	}
}
