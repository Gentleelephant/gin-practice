package v1

import (
	"gin-practice/src/global"
	"gin-practice/src/model"
	"gin-practice/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
func Register(c *gin.Context) {

	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	dto := &model.UserRegisterDTO{}
	err = c.ShouldBindJSON(dto)
	if err != nil {
		return
	}
	encryption, err := utils.Encryption(dto.Password)
	if err != nil {
		return
	}
	dto.Password = encryption
	err = global.UserDao.CreateUser(model.UserRegisterDTOToUser(dto))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.CustomResp{
		Code: 2000,
		Msg:  "注册成功",
		Data: map[string]string{
			"username": dto.Username,
			"email":    dto.Email,
			"phone":    dto.Phone,
		},
	})
}
