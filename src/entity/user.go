package entity

type User struct {
	Username string `json:"username"`

	Password string `json:"password"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}

type LoginUser struct {
	Username string `json:"username" form:"username" binding:"required"`

	Password string `json:"password" form:"password" binding:"required"`
}
