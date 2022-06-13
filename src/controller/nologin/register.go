package nologin

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gin-practice/src/model"
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

	dto := &model.UserDTO{}
	err = c.BindJSON(dto)
	if err != nil {
		return
	}
	db := config.DB
	err = db.Model(&entity.User{}).Create(model.UserDTOToUser(dto)).Error
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
