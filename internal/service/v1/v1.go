package service

import "context"

type stream interface {
	Create(context.Context, StreamConnection)
	Watch(context.Context, StreamConnection)
}

type Service struct {
	Stream stream
}

func NewService() *Service {
	return &Service{
		Stream: newStreamService(),
	}
}
