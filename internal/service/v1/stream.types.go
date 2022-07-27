package service

import "github.com/pion/webrtc/v3"

type StreamConnection interface {
	Read() (*StreamMessage, error)
	Write(*StreamMessage) error
	WriteError(error) error
}

type StreamMessage struct {
	BroadcastID string
	Session     *webrtc.SessionDescription
}

func NewStreamMessage(broadcastID, sdpType, sdp string) *StreamMessage {
	return &StreamMessage{
		BroadcastID: broadcastID,
		Session: &webrtc.SessionDescription{
			Type: webrtc.NewSDPType(sdpType),
			SDP:  sdp,
		},
	}
}
