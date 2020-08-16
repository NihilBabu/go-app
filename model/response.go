package model

type Response struct {
	Status  bool
	Message string
	Data    string
}

func SuccessResponse(data string) Response {
	return Response{
		Status:  true,
		Message: "",
		Data:    data,
	}
}

func SuccessResponse(message, data string) Response {
	return Response{
		Status:  true,
		Message: "",
		Data:    data,
	}
}
