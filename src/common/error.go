package common

type CommonError struct {
	Code int `json:"code"`

	Msg string `json:"msg"`
}
