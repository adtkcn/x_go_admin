package system

import (
	"x-gin-admin/utils/captcha"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct{}

func (u *CaptchaController) Generate(c *gin.Context) {
	id, b64s, err := captcha.Generate()
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", gin.H{"id": id, "b64s": b64s})
}

// func (u *CaptchaController) Verify(c *gin.Context) {
// 	var verifyID = c.Query("id")
// 	var verifyCode = c.Query("code")
// 	isPass := captcha.Verify(verifyID, verifyCode)

// 	response.Send(c, "ok", gin.H{"isPass": isPass})
// }
