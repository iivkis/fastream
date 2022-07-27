package service

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
)

type chatService struct {
	mx          sync.Mutex
	subscribers map[chatSubscriberUUID]chan *ChatMessage
}

func newChatService() chat {
	return &chatService{
		subscribers: make(map[string]chan *ChatMessage),
	}
}

func (c *chatService) subscribe() (chatSubscriberUUID, chan *ChatMessage) {
	id, ch := uuid.NewString(), make(chan *ChatMessage, 1)
	c.subscribers[id] = ch
	return id, ch
}

func (c *chatService) unsubscride(id chatSubscriberUUID) {
	c.mx.Lock()
	defer c.mx.Unlock()

	close(c.subscribers[id])
	delete(c.subscribers, id)
}

func (c *chatService) sendToSubsribers(message *ChatMessage) {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, subscriber := range c.subscribers {
		subscriber <- message
	}
}

func (c *chatService) Connect(ctx context.Context, conn ChatConnection) {
	go func() {
		id, subscribe := c.subscribe()
		defer c.unsubscride(id)

		for {
			select {
			case message := <-subscribe:
				if err := conn.Write(message); err != nil {
					log.Println(err)
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		receivedMessages, err := conn.Read()
		if err != nil {
			log.Println(err)
			return
		}

		c.sendToSubsribers(receivedMessages)
	}
}
