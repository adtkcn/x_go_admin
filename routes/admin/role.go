package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseRole(r *gin.RouterGroup) {
	role := system.RoleController{}
	r.GET("role/List", role.List)
	r.POST("role/Create", role.Create)
	r.POST("role/Update", role.Update)
	r.POST("role/Delete", role.Delete)

	r.GET("role/FindOne", role.FindOne)
}
