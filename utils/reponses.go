package utils

type Response struct {
	IsError    bool        `json:"is_error"`
	ErrMessage *string     `json:"err_message"`
	Data       interface{} `json:"data,omitempty"`
}

func ErrorResp(message string) Response {
	return Response{
		IsError:    true,
		ErrMessage: &message,
	}
}

func SuccessResp(data ...interface{}) Response {
	var responseData interface{} = nil
	if len(data) > 0 {
		responseData = data[0]
	}
	return Response{
		IsError:    false,
		ErrMessage: nil,
		Data:       responseData,
	}
}
