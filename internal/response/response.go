package response

type Response struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

type ValidationError struct {
	Field string      `json:"field"`
	Error string      `json:"error"`
	Value interface{} `json:"value,omitempty"`
}

func OK() Response {
	return Response{
		Status: "success",
	}
}

func Success(data interface{}) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

func Error(message string, data interface{}) Response {
	res := Response{
		Status:  "error",
		Message: message,
		Data:    data,
	}
	return res
}
