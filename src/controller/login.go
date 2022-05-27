package controller

import (
	"gin-practice/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	loginUser := &entity.LoginUser{}

	err := c.Bind(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4001,
			Msg:  "参数错误",
			Data: err.Error(),
		})
		c.Abort()
		return
	}

	c.SetCookie("token", "token", 120, "/", "127.0.0.1", false, true)

	c.JSON(http.StatusOK, entity.CustomResp{
		Code: 2000,
		Msg:  "登录成功",
		Data: loginUser,
	})

}
