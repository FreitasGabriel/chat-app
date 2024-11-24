package entity

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewMessage(username, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
	}
}
