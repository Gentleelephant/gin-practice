package controller

import (
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

	c.SetCookie("token", "token", 120, "/", c.GetHeader("Host"), false, true)

	c.JSON(http.StatusOK, entity.CustomResp{
		Code: 2000,
		Msg:  "登录成功",
		Data: loginUser,
	})

}
