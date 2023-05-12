package response

import "github.com/gin-gonic/gin"

type QueryRes struct {
	Count int64 `json:"count"`
	List  any   `json:"list"`
	// Page  int   `json:"page"`
}
type send struct {
	code int
	msg  string
	data QueryRes
}

func Send(c *gin.Context, msg string, data any) {
	var res = &gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	}
	c.JSON(200, res)
}
func SendList(c *gin.Context, msg string, count int64, list []any) {
	var res = &send{
		code: 200,
		msg:  msg,
		data: QueryRes{
			Count: count,
			List:  &list,
		},
	}
	c.JSON(200, res)
}

func SendError(c *gin.Context, msg string, data any) {
	var res = &gin.H{
		"code": 500,
		"msg":  msg,
		"data": data,
	}
	c.JSON(200, res)
}
