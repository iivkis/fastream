package restfulv1

type stdResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

func newResponse(data interface{}, err error) *stdResponse {
	resp := stdResponse{Data: data}
	if err != nil {
		resp.Error = err.Error()
	}
	return &resp
}
