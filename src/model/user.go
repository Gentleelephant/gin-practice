package model

type UserRegisterDTO struct {
	Username string `json:"username" binding:"required"`

	Password string `json:"password" binding:"required"`

	Email string `json:"email" binding:"required"`

	Phone string `json:"phone"`
}

type LoginUser struct {
	Username string `json:"username" form:"username" binding:"required"`

	Password string `json:"password" form:"password" binding:"required"`
}

type UserRegisterVO struct {
	Username string `json:"username"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}
