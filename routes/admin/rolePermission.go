package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseRolePermission(r *gin.RouterGroup) {
	rolePermission := system.RolePermissionController{}
	// r.GET("rolePermission/List", rolePermission.List)
	r.POST("rolePermission/Create", rolePermission.Create)
	// r.POST("rolePermission/Update", rolePermission.Update)
	r.POST("rolePermission/Delete", rolePermission.Delete)
	// r.GET("rolePermission/FindOne", rolePermission.FindOne)
}
