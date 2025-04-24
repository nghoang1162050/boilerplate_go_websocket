package dto

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBaseResponse(code int, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
