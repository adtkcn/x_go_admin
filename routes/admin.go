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

	captcha := system.CaptchaController{}
	r.GET("captcha/Generate", captcha.Generate)

	common := system.CommonController{}
	r.POST("common/upload", handler.CalculateFileMD5(), common.Upload)
	r.GET("common/FileByID", common.FileByID)
	r.GET("common/CopyWithMd5", common.CopyWithMd5)

	menu := system.MenuController{}
	permission := system.PermissionController{}
	/**
	* 子级路由组
	* server/admin/
	 */
	admin := r.Group("/admin", handler.GetUserId())
	{
		// handler.GetUserInfo()

		admin.GET("user/GetUsers", user.GetUsers)
		admin.POST("user/Logout", user.Logout)
		admin.POST("user/UpdateUser", user.UpdateUser)
		admin.POST("user/DeleteUser", user.DeleteUser)

		role := system.RoleController{}
		admin.GET("role/List", role.List)

		admin.GET("menu/List", menu.List)
		admin.POST("menu/Create", menu.Create)
		admin.POST("menu/Update", menu.Update)
		admin.POST("menu/Delete", menu.Delete)
		admin.GET("menu/GetMenusByUser", menu.GetMenusByUser)

		admin.GET("permission/GetUserPermission", permission.GetUserPermission)
		// admin.GET("captcha/Verify", captcha.Verify)
	}
}
