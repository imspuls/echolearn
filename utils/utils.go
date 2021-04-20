package utils

type (
	ResponseBase struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	Response struct {
		ResponseBase
		Data interface{} `json:"data"`
	}
)

func Message(code int, message string) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": message}
}

func Respond(code int, message string) ResponseBase {
	var resp ResponseBase
	resp.Code = code
	resp.Message = message
	return resp
}

func RespondWithData(code int, message string, data interface{}) Response {
	var resp Response
	resp.Code = code
	resp.Message = message
	resp.Data = data
	return resp
}
