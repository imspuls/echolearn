package models

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

var (
	res Response
)
