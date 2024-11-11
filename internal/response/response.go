package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
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

func Error(message string, code ...int) Response {
	res := Response{
		Status:  "error",
		Message: message,
	}
	if len(code) > 0 {
		res.Code = code[0]
	}
	return res
}
