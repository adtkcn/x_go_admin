package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(c *gin.Context, msg string, data any) {
	var res = &gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	}
	c.JSON(http.StatusOK, res)
}

func SendError(c *gin.Context, msg string, data any) {
	var res = &gin.H{
		"code": 500,
		"msg":  msg,
		"data": data,
	}
	c.JSON(http.StatusOK, res)
}
