package v1

import (
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// Hello PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 2000 {string} hello world
// @Router /hello [get]
func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, model.CustomResp{
		Code: 2000,
		Msg:  "hello world",
		Data: "",
	})

}
