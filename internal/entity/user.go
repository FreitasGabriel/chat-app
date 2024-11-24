package entity

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(name, email, username, password string) *User {
	return &User{
		ID:       NewUUID().String(),
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
}
