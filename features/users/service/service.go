package service

import (
	"boilerplate-feature/features/users"
	"fmt"
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
	err := u.validate.Struct(Input)
	if err != nil {
		return err
	}
	if len(Input.Password) < 8 {
		return fmt.Errorf("validation error password must be at least 8 characters")
	}
	if Input.Email == "" {
		return fmt.Errorf("validation error email cannot be empty")
	}
	if Input.Phone == "" {
		return fmt.Errorf("validation error phone cannot be empty")
	}
	if Input.Role == "" {
		return fmt.Errorf("validation error role cannot be empty")
	}
	if Input.Username == "" {
		return fmt.Errorf("validation error username cannot be empty")
	}
	if Input.FirstName == "" || Input.LastName == "" {
		return fmt.Errorf("validation error first name and last name cannot be empty")
	}
	errRegister := u.userData.Register(Input)
	if errRegister != nil {
		return errRegister
	}
	return nil
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
