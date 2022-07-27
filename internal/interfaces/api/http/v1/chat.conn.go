package restful

import (
	"github.com/gorilla/websocket"
	"github.com/iivkis/fastream/internal/service/v1"
)

type chatConnection struct {
	ws *websocket.Conn
}

func newChatConnection(conn *websocket.Conn) service.ChatConnection {
	return &chatConnection{conn}
}

func (c *chatConnection) Read() (*service.ChatMessage, error) {
	var d chatMessage

	if err := c.ws.ReadJSON(&d); err != nil {
		return nil, err
	}

	m := service.NewChatMessage(d.Username, d.Content)
	return m, nil
}

func (c *chatConnection) Write(m *service.ChatMessage) error {
	d := newChatMessage(m.Username, m.Content)
	return c.ws.WriteJSON(NewResponse(d, nil))
}

func (c *chatConnection) WriteError(err error) error {
	return c.ws.WriteJSON(NewResponse(nil, err))
}
