package model

import (
	"gin-practice/src/entity"
)

type UserDTO struct {
	Username string `json:"username" binding:"required"`

	Password string `json:"password" binding:"required"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}

type LoginUser struct {
	Username string `json:"username" form:"username" binding:"required"`

	Password string `json:"password" form:"password" binding:"required"`
}

// UserVO TODO
type UserVO struct {
	Username string `json:"username"`

	Email string `json:"email"`

	Phone string `json:"phone"`
}

func UserDTOToUser(dto *UserDTO) *entity.User {
	return &entity.User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Phone:    dto.Phone,
	}
}
