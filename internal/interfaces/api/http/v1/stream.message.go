package restful

import "github.com/pion/webrtc/v3"

type streamMessage struct {
	BroadcastID string `json:"broadcastID,omitempty"`
	Type        string `json:"type,omitempty"`
	SDP         string `json:"sdp,omitempty"`
}

func newStreamMessage(broadcastID string, sdpType webrtc.SDPType, sdp string) *streamMessage {
	return &streamMessage{
		BroadcastID: broadcastID,
		Type:        sdpType.String(),
		SDP:         sdp,
	}
}
