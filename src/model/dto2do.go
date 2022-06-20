package model

import "gin-practice/src/entity"

func UserRegisterDTOToUser(dto *UserRegisterDTO) *entity.User {
	return &entity.User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Phone:    dto.Phone,
	}
}
