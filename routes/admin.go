package routes

import (
	"x-gin-admin/routes/admin"
	"x-gin-admin/utils/handler"

	"github.com/gin-gonic/gin"
)

func UseAdmin(r *gin.RouterGroup) {
	/**
	* 子级路由组
	* server/admin/
	 */
	adminRoute := r.Group("/admin", handler.GetUserId())
	admin.UseUser(adminRoute)
	admin.UseUserRole(adminRoute)

	admin.UseMenu(adminRoute)
	admin.UsePermission(adminRoute)

	admin.UseRole(adminRoute)
	admin.UseRoleMenu(adminRoute)
	admin.UseRolePermission(adminRoute)

}
