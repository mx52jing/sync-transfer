package utils

type CommonResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}