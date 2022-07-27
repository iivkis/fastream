package restful

import "encoding/json"

type StandartResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(data interface{}, err error) *StandartResponse {
	resp := StandartResponse{Data: data}
	if err != nil {
		resp.Error = err.Error()
	}
	return &resp
}

func (r *StandartResponse) JSON() ([]byte, error) {
	return json.Marshal(r)
}
