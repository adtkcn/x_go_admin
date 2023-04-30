package response

import "github.com/gin-gonic/gin"

func Send(c *gin.Context, message string, data any) {
	var res = &gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	}
	c.JSON(200, res)
}
func SendError(c *gin.Context, message string, data any) {
	var res = &gin.H{
		"code":    500,
		"message": message,
		"data":    data,
	}
	c.JSON(500, res)
}
