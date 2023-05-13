package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type RoleMenuController struct{}

// 添加角色的菜单
func (u *RoleMenuController) Create(c *gin.Context) {
	var rolePermission model.RoleMenu
	if err := c.ShouldBind(&rolePermission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	if rolePermission.MenuID == 0 || rolePermission.RoleID == 0 {
		response.SendError(c, "未指定菜单或者角色", nil)
		return
	}
	err := db.Sql.Create(&rolePermission).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &rolePermission)
}

// 删除角色的菜单
func (u *RoleMenuController) Delete(c *gin.Context) {
	var rolePermission model.RoleMenu
	if err := c.ShouldBind(&rolePermission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	if rolePermission.MenuID == 0 || rolePermission.RoleID == 0 {
		response.SendError(c, "未指定菜单或者角色", nil)
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
