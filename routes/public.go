package routes

import (
	"x-gin-admin/admin/controller/system"
	"x-gin-admin/utils/handler"

	"github.com/gin-gonic/gin"
)

func UsePublic(r *gin.RouterGroup) {
	user := system.UserController{}
	r.POST("user/Register", user.Register)
	r.POST("user/Login", user.Login)

	captcha := system.CaptchaController{}
	r.GET("captcha/Generate", captcha.Generate)

	file := system.FileController{}
	r.POST("file/Upload", handler.CalculateFileMD5(), file.Upload)
	r.GET("file/FileByID", file.FileByID)
	r.GET("file/UploadWithMd5", file.UploadWithMd5)
}
