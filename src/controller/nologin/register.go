package nologin

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 注册用户
func Register(c *gin.Context) {

	dto := &model.UserDTO{}
	err := c.BindJSON(dto)
	if err != nil {
		c.JSON(http.StatusOK, model.CustomResp{
			Code: 4001,
			Msg:  "参数错误",
			Data: err.Error(),
		})
		c.Abort()
		return
	}

	db := config.DB
	if err != nil {
		log.Println(err)
	}
	db.Model(&entity.User{}).Create(model.UserDTOToUser(dto))
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
