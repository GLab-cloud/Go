package model

type Response struct {
	StatusCode int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty" db:"	"` //object
}

func NewResponse(statusCode int, message string, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
