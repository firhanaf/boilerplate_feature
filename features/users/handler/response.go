package handler

import "boilerplate-feature/features/users"

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func UserCoretoResponse(input *users.UserCore) UserResponse {
	return UserResponse{
		ID:        input.ID,
		Username:  input.Username,
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Role:      input.Role,
	}
}
