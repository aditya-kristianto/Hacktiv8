package views

type Response struct {
	StatusCode int         `json:"status" example:"200"`
	Message    string      `json:"message" example:"GET_SUCCESS"`
	Payload    interface{} `json:"payload"`
}
