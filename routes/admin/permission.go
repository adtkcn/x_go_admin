package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UsePermission(r *gin.RouterGroup) {
	permission := system.PermissionController{}
	r.GET("permission/GetUserPermission", permission.GetUserPermission)
	r.POST("permission/Create", permission.Create)
}
