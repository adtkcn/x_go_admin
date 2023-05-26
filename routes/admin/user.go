package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseUser(r *gin.RouterGroup) {
	user := system.UserController{}
	r.GET("user/GetUsers", user.GetUsers)
	r.POST("user/Logout", user.Logout)

	r.GET("user/List", user.List)
	r.POST("user/Create", user.Create)
	r.POST("user/Update", user.Update)
	r.POST("user/Delete", user.Delete)
}
