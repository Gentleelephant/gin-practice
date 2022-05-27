package controller

import (
	"gin-practice/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, entity.CustomResp{
		Code: 2000,
		Msg:  "hello world",
		Data: "",
	})

}
