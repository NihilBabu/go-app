package dto

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

func SuccessResponseWithMessage(data,message string) Response {
	return Response{
		Status:  true,
		Message: "",
		Data:    data,
	}
}
