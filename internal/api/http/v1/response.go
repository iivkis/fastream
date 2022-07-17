package apihv1

type stdResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func newResponse(data interface{}, err error) *stdResponse {
	reponse := stdResponse{}
	reponse.Data = data

	if err != nil {
		reponse.Error = err.Error()
	}

	return &reponse
}
