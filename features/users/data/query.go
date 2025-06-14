package data

import (
	"boilerplate-feature/features/users"
	"boilerplate-feature/helpers"
	"errors"
	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func (repo *UserQuery) Login(Identifier string, Password string) (*users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserQuery) UpdateProfile(Input users.UserCore) error {
	//TODO implement me
	panic("implement me")
}

func (repo *UserQuery) GetProfile(ID string) (*users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserQuery) GetAllUsers() ([]users.UserCore, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserQuery) DeleteAccount(ID string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *UserQuery) Register(user users.UserCore) error {
	var userModel = UserCoretoModel(user)
	nanoid, errID := helpers.NanoIDGenerator()
	if errID != nil {
		return errID
	}
	userModel.ID = nanoid

	hash, err := helpers.HashPassword(userModel.Password)
	if err != nil {
		return err
	}
	userModel.Password = hash
	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("no row affected")
	}
	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}
