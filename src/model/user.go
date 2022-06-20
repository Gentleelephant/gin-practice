package model

// UserRegisterDTO 用户注册请求传输对象
type UserRegisterDTO struct {
	Username string `json:"username" binding:"required"`

	Password string `json:"password" binding:"required"`

	Email string `json:"email" binding:"required"`

	Phone string `json:"phone"`
}

// LoginUserDTO 登录请求传输对象
type LoginUserDTO struct {
	Username string `json:"username" form:"username" binding:"required"`

	Password string `json:"password" form:"password" binding:"required"`
}

// UserRegisterVO 用户注册成功响应数据对象
type UserRegisterVO struct {
	Username string `json:"username"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}
