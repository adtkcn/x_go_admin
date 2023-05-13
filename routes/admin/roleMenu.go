package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseRoleMenu(r *gin.RouterGroup) {
	roleMenuController := system.RoleMenuController{}

	r.POST("roleMenu/Create", roleMenuController.Create)
	r.POST("roleMenu/Delete", roleMenuController.Delete)

}
