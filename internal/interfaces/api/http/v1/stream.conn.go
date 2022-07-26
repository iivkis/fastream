package restfulv1

import (
	"github.com/gorilla/websocket"
	servicev1 "github.com/iivkis/fastream/internal/service/v1"
	"github.com/pion/webrtc/v3"
)

type streamConnectionMessage struct {
	BroadcastID string `json:"broadcastID,omitempty"`
	Type        string `json:"type,omitempty"`
	SDP         string `json:"sdp,omitempty"`
}

func newStreamConnectionMessage(broadcastID string, sdpType webrtc.SDPType, sdp string) *streamConnectionMessage {
	return &streamConnectionMessage{
		BroadcastID: broadcastID,
		Type:        sdpType.String(),
		SDP:         sdp,
	}
}

type streamConnection struct {
	*websocket.Conn
}

func newStreamConnection(c *websocket.Conn) servicev1.StreamConnection {
	return &streamConnection{c}
}

func (c *streamConnection) Read() (*servicev1.StreamMessage, error) {
	var d streamConnectionMessage

	err := c.ReadJSON(&d)
	if err != nil {
		return nil, err
	}

	m := servicev1.NewStreamMessage(d.BroadcastID, d.Type, d.SDP)
	return m, nil
}

func (c *streamConnection) Write(m *servicev1.StreamMessage) error {
	d := newStreamConnectionMessage(
		m.BroadcastID,
		m.Session.Type,
		m.Session.SDP,
	)

	return c.WriteJSON(NewResponse(d, nil))
}

func (c *streamConnection) WriteError(e error) error {
	return c.WriteJSON(NewResponse(nil, e))
}
