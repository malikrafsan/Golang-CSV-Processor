package contracts

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailedResponse(message string, status int) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  200,
		Message: message,
		Data:    data,
	}
}
