package handler

import "boilerplate-feature/features/users"

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,numeric"`
	Role      string `json:"role"`
}

func UserRequesttoCore(input UserRequest) users.UserCore {
	return users.UserCore{
		Username:  input.Username,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Role:      input.Role,
	}
}
