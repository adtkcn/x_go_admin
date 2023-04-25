package routes

import (
	"x-gin-admin/admin/controller/system"
	"x-gin-admin/utils/handler"

	"github.com/gin-gonic/gin"
)

func Use(r *gin.RouterGroup) {
	/**
	* 子级路由组
	* server/admin/
	 */
	admin := r.Group("/admin", handler.GetUserId(), handler.GetUserInfo())
	{
		user := system.UserController{}
		admin.GET("user/GetUsers", user.GetUsers)
		admin.GET("user/Register", user.Register)
		admin.GET("user/Login", user.Login)
		admin.GET("user/UpdateUser", user.UpdateUser)
		admin.GET("user/DeleteUser", user.DeleteUser)

		role := system.RoleController{}
		admin.GET("role/GetRoles", role.GetRoles)
	}

	// api := r.Group("/api")
	{
		// var user = controller.User{}
		// api.GET("", user.SelectUser)
	}
}
