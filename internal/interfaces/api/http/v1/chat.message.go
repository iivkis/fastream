package restful

type chatMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func newChatMessage(username, content string) *chatMessage {
	return &chatMessage{Username: username, Content: content}
}
