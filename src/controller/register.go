package controller

import (
	"gin-practice/src/entity"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
func Register(c *gin.Context) {

	dto := &model.UserDTO{}
	err := c.BindJSON(dto)
	if err != nil {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4001,
			Msg:  "参数错误",
			Data: err.Error(),
		})
		c.Abort()
		return
	}

}
