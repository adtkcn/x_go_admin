package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseUserRole(r *gin.RouterGroup) {
	userRole := system.UserRoleController{}
	r.GET("userRole/GetRoleByUserId", userRole.GetRoleByUserId)
	r.GET("userRole/GetUserPermissionByUserId", userRole.GetUserPermissionByUserId)

	r.POST("userRole/Create", userRole.Create)
	// r.POST("userRole/Update", userRole.Update)
	r.POST("userRole/Delete", userRole.Delete)

}
