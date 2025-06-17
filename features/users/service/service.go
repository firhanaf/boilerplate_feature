package service

import (
	"boilerplate-feature/app/middlewares"
	"boilerplate-feature/features/users"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func (u UserService) Login(Identifier string, Password string) (string, error) {
	login, errLogin := u.userData.Login(Identifier, Password)
	if errLogin != nil {
		return "", errLogin
	}
	token, err := middlewares.CreateToken(login.ID)
	if err != nil {
		return "", err
	}
	return token, nil

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

func (u UserService) UpdateProfile(ID string, input users.UserCore) error {
	// Validasi isi struct, tapi gunakan tag validate:"omitempty"
	err := u.validate.Struct(input)
	if err != nil {
		return fmt.Errorf("validation error: %s", err.Error())
	}

	// Validasi khusus password hanya jika dikirim
	if input.Password != "" && len(input.Password) < 8 {
		return fmt.Errorf("validation error: password must be at least 8 characters")
	}

	// Lanjut update ke data layer
	errUpdate := u.userData.UpdateProfile(ID, input)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (u UserService) GetProfile(ID string) (users.UserCore, error) {
	result, err := u.userData.GetProfile(ID)
	if err != nil {
		return users.UserCore{}, err
	}
	return result, nil
}

func (u UserService) GetAllUsers(ID string) ([]users.UserCore, error) {
	result, err := u.userData.GetAllUsers(ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserService) DeleteAccount(ID string) error {
	err := u.userData.DeleteAccount(ID)
	if err != nil {
		return err
	}
	return nil
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
		validate: validator.New(),
	}
}
