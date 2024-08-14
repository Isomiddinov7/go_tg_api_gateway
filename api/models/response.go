package models

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ResponseModelRespondent struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
