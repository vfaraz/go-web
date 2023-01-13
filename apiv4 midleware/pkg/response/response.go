package response

type Response struct {
	Message string
	Data    interface{} //any
}

func Ok(message string, data interface{}) Response {
	return Response{Message: message, Data: data}
}

func Error(err error) Response {
	return Response{Message: err.Error(), Data: nil}
}
