package apihv1

type StandartResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func NewResponse(data interface{}, err error) *StandartResponse {
	reponse := StandartResponse{}
	reponse.Data = data

	if err != nil {
		reponse.Error = err.Error()
	}

	return &reponse
}
