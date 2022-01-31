package responses

type webResponse struct {
	Code   int         `json:"code"`
	Message string      `json:"message"`
	Data   interface{} `json:"data"`
}


func GenerateResponse(statusCode int, msg string, data interface{}) webResponse {
	return webResponse{
		Code:   statusCode,
		Message: msg,
		Data:   data,
	}
}