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
	var user User
	tx := repo.db.Where("email = ? or phone = ? or username = ?", Identifier, Identifier, Identifier).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	checkPass := helpers.CheckPasswordHash(Password, user.Password)
	if !checkPass {
		return nil, errors.New("invalid identifier or password")
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("invalid identifier or password")
	}

	return &users.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func (repo *UserQuery) UpdateProfile(ID string, input users.UserCore) error {
	var existingUser User
	tx := repo.db.First(&existingUser, "id = ?", ID)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return tx.Error
	}

	// Siapkan data yang ingin diupdate secara dinamis
	updateData := map[string]interface{}{}

	if input.Username != "" {
		updateData["username"] = input.Username
	}
	if input.Email != "" {
		updateData["email"] = input.Email
	}
	if input.Phone != "" {
		updateData["phone"] = input.Phone
	}
	if input.FirstName != "" {
		updateData["first_name"] = input.FirstName
	}
	if input.LastName != "" {
		updateData["last_name"] = input.LastName
	}
	if input.Role != "" {
		updateData["role"] = input.Role
	}
	if input.Password != "" {
		// Hash password hanya jika password diberikan
		hashed, err := helpers.HashPassword(input.Password)
		if err != nil {
			return err
		}
		updateData["password"] = hashed
	}

	if len(updateData) == 0 {
		return errors.New("no valid fields to update")
	}

	// Update hanya field yang dikirim
	tx = repo.db.Model(&existingUser).Updates(updateData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *UserQuery) GetProfile(ID string) (users.UserCore, error) {
	var user User
	tx := repo.db.First(&user, "id = ?", ID)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected != 1 {
		return users.UserCore{}, errors.New("user not found")
	}
	result := users.UserCore{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
	}
	return result, nil
}

func (repo *UserQuery) GetAllUsers(ID string) ([]users.UserCore, error) {
	var user []User
	tx := repo.db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var userCore []users.UserCore
	for _, v := range user {
		userCore = append(userCore, users.UserCore{
			ID:        v.ID,
			Username:  v.Username,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Phone:     v.Phone,
			Role:      v.Role,
		})
	}
	return userCore, nil
}

func (repo *UserQuery) DeleteAccount(ID string) error {
	var user User
	tx := repo.db.First(&user, "id = ?", ID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("user not found")
	}
	tx = repo.db.Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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
