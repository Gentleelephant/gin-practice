package common

// CommonError is the error struct for common error

type CommonError struct {
	Code int `json:"code"`

	Msg string `json:"msg"`
}

var (
	ServerError = &CommonError{
		Code: 500,
		Msg:  "server error",
	}
)

func (e CommonError) Error() string {
	return e.Msg
}
