package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseUser(r *gin.RouterGroup) {
	user := system.UserController{}
	r.GET("user/GetUsers", user.GetUsers)
	r.POST("user/Logout", user.Logout)
	r.POST("user/UpdateUser", user.UpdateUser)
	r.POST("user/DeleteUser", user.DeleteUser)
}
