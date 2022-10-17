package helper

type (
	Response struct {
		Status  int         `json:"status,omitempty"`
		Message string      `json:"message"`
		Error   interface{} `json:"error,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	}
)
