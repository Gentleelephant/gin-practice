package controller

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	loginUser := &model.LoginUser{}

	err := c.BindJSON(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4001,
			Msg:  "参数错误",
			Data: err.Error(),
		})
		c.Abort()
		return
	}

	db := config.GolbalConfig.DB

	dbUser := &entity.User{}

	db.Where(&entity.User{Username: loginUser.Username}, "username").Find(&dbUser)

	if dbUser.Username == "" {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4002,
			Msg:  "用户不存在",
			Data: "",
		})
		c.Abort()
		return
	}

	if dbUser.Password != loginUser.Password {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4003,
			Msg:  "密码错误",
			Data: "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, entity.CustomResp{
		Code: 2000,
		Msg:  "登录成功",
		Data: "",
	})

	c.SetCookie("token", "token", 120, "/", c.GetHeader("Host"), false, true)

}
