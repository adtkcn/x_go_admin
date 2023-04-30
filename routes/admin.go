package routes

import (
	"x-gin-admin/admin/controller/system"
	"x-gin-admin/utils/handler"

	"github.com/gin-gonic/gin"
)

func UseAdmin(r *gin.RouterGroup) {
	user := system.UserController{}

	r.POST("user/Register", user.Register)
	r.POST("user/Login", user.Login)

	/**
	* 子级路由组
	* server/admin/
	 */
	admin := r.Group("/admin", handler.GetUserId())
	{
		// handler.GetUserInfo()
		admin.GET("user/GetUsers", user.GetUsers)

		admin.POST("user/UpdateUser", user.UpdateUser)
		admin.POST("user/DeleteUser", user.DeleteUser)

		role := system.RoleController{}
		admin.GET("role/GetRoles", role.GetRoles)

		captcha := system.CaptchaController{}
		admin.GET("captcha/Generate", captcha.Generate)
		// admin.GET("captcha/Verify", captcha.Verify)
	}
}
