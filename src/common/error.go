package common

// CommonError is the error struct for common error

type CommonError struct {
	Code int `json:"code"`

	Msg string `json:"msg"`

	MsgCn string `json:"msg_cn"`
}

var (
	InvalidParamsError = &CommonError{
		Code:  400,
		Msg:   "invalid params",
		MsgCn: "参数错误",
	}

	UnauthorizedError = &CommonError{
		Code:  401,
		Msg:   "unauthorized",
		MsgCn: "未授权",
	}

	ForbiddenError = &CommonError{
		Code:  403,
		Msg:   "forbidden",
		MsgCn: "禁止访问",
	}

	InternalServerError = &CommonError{
		Code:  500,
		Msg:   "internal server error",
		MsgCn: "内部服务器错误",
	}

	UserNotFoundError = &CommonError{
		Code:  404,
		Msg:   "user not found",
		MsgCn: "用户未找到",
	}

	UsernameAlreadyExistError = &CommonError{
		Code:  400,
		Msg:   "username already exist",
		MsgCn: "用户名已存在",
	}

	EmailAlreadyExistError = &CommonError{
		Code:  400,
		Msg:   "email already exist",
		MsgCn: "邮箱已存在",
	}

	PasswordWrongError = &CommonError{
		Code:  400,
		Msg:   "password wrong",
		MsgCn: "密码错误",
	}
)

func (e CommonError) Error() string {
	return e.Msg
}
