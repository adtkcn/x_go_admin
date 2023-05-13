package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type RolePermissionController struct{}

// 添加用户角色
func (u *RolePermissionController) Create(c *gin.Context) {
	var rolePermission model.RolePermission
	if err := c.ShouldBind(&rolePermission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	if rolePermission.PermissionID == 0 || rolePermission.RoleID == 0 {
		response.SendError(c, "未指定权限或者角色", nil)
		return
	}
	err := db.Sql.Create(&rolePermission).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &rolePermission)
}
func (u *RolePermissionController) Delete(c *gin.Context) {
	var rolePermission model.RolePermission
	if err := c.ShouldBind(&rolePermission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	if rolePermission.PermissionID == 0 || rolePermission.RoleID == 0 {
		response.SendError(c, "未指定权限或者角色", nil)
		return
	}
	// err := roleService.DeleteById(id)
	err := db.Sql.Delete(&rolePermission).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "删除成功", nil)
}
