package dto

type ChangePasswordDTO struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type CreateUserOutput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FindUserOutput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
