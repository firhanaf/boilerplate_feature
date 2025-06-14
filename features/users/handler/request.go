package handler

import "boilerplate-feature/features/users"

type UserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,numeric"`
	Role      string `json:"role"`
}

type UserUpdateRequest struct {
	Username  string `json:"username" validate:"omitempty,min=3"`
	Email     string `json:"email" validate:"omitempty,email"`
	Phone     string `json:"phone" validate:"omitempty"`
	FirstName string `json:"firstname" validate:"omitempty"`
	LastName  string `json:"lastname" validate:"omitempty"`
	Password  string `json:"password" validate:"omitempty,min=8"`
	Role      string `json:"role" validate:"omitempty"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
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

func UserUpdateRequesttoCore(input UserUpdateRequest) users.UserCore {
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
