package entity

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewUser(name, email, username, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
}

func NewMessage(username, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
	}
}
