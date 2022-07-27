package service

import "context"

type stream interface {
	Create(context.Context, StreamConnection)
	Watch(context.Context, StreamConnection)
}

type utils interface {
	GetLocalIP() (localIP string, err error)
}

type Service struct {
	Stream stream
	Utils  utils
}

func NewService() *Service {
	return &Service{
		Stream: newStreamService(),
		Utils:  newUtilsService(),
	}
}
