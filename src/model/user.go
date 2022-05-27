package model

type UserDTO struct {
	Username string `json:"username"`

	Password string `json:"password"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}

// UserVO TODO
type UserVO struct {
}
