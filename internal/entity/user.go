package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(name, email, username, password string) *User {

	hash, _ := GenerateHashedPassword(password)

	return &User{
		ID:       NewUUID().String(),
		Name:     name,
		Email:    email,
		Username: username,
		Password: string(hash),
	}
}

func GenerateHashedPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (u *User) ValidatePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil, err
}
