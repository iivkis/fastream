package service

type ChatConnection interface {
	Read() (*ChatMessage, error)
	Write(*ChatMessage) error
	WriteError(error) error
}

type chatSubscriberUUID = string

type ChatMessage struct {
	Username string
	Content  string
}

func NewChatMessage(username, content string) *ChatMessage {
	return &ChatMessage{
		Username: username,
		Content:  content,
	}
}
