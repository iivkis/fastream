package service

import "context"

type stream interface {
	Create(context.Context, StreamConnection)
	Watch(context.Context, StreamConnection)
}

type utils interface {
	GetLocalIP() (localIP string, err error)
}

type chat interface {
	Connect(context.Context, ChatConnection)
}

type Service struct {
	Stream stream
	Utils  utils
	Chat   chat
}

func NewService() *Service {
	return &Service{
		Stream: newStreamService(),
		Utils:  newUtilsService(),
		Chat:   newChatService(),
	}
}
