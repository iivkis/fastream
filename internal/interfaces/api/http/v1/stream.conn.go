package restful

import (
	"github.com/gorilla/websocket"
	"github.com/iivkis/fastream/internal/service/v1"
)

type streamConn struct {
	*websocket.Conn
}

func newStreamConn(c *websocket.Conn) service.StreamConnection {
	return &streamConn{c}
}

func (c *streamConn) Read() (*service.StreamMessage, error) {
	var d streamMessage

	err := c.ReadJSON(&d)
	if err != nil {
		return nil, err
	}

	m := service.NewStreamMessage(d.BroadcastID, d.Type, d.SDP)
	return m, nil
}

func (c *streamConn) Write(m *service.StreamMessage) error {
	d := newStreamMessage(
		m.BroadcastID,
		m.Session.Type,
		m.Session.SDP,
	)

	return c.WriteJSON(NewResponse(d, nil))
}

func (c *streamConn) WriteError(e error) error {
	return c.WriteJSON(NewResponse(nil, e))
}
