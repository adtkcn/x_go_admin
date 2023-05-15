package system

import (
	"x-gin-admin/admin/service/system"
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

var userRoleService = system.UserRoleService{}

type UserRoleController struct{}

// 添加用户角色
func (u *UserRoleController) Create(c *gin.Context) {
	var userRole model.UserRole
	if err := c.ShouldBind(&userRole); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	if userRole.UserID == 0 || userRole.RoleID == 0 {
		response.SendError(c, "未指定用户或者角色", nil)
		return
	}
	err := db.Sql.Create(&userRole).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &userRole)
}

// 删除用户角色
func (u *UserRoleController) Delete(c *gin.Context) {
	var userRole model.UserRole
	if err := c.ShouldBind(&userRole); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	if userRole.UserID == 0 || userRole.RoleID == 0 {
		response.SendError(c, "未指定用户或者角色", nil)
		return
	}
	// err := roleService.DeleteById(id)
	err := db.Sql.Delete(&userRole).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "删除成功", nil)
}

// 获取用户的角色
func (u *UserRoleController) GetRoleByUserId(c *gin.Context) {
	var userRole model.UserRole
	if err := c.ShouldBindQuery(&userRole); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	if userRole.UserID == 0 {
		response.SendError(c, "未指定用户", nil)
		return
	}
	list, _, err := userRoleService.FindByUserId(userRole.UserID)

	// var list []model.UserRole
	// err := db.Sql.Model(&model.Role{}).Where("role_id = ?", userRole.UserID).Find(&list).Error

	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &list)
}

// 获取用户的角色和权限
func (u *UserRoleController) GetUserPermissionByUserId(c *gin.Context) {
	var userRole model.UserRole
	if err := c.ShouldBindQuery(&userRole); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	if userRole.UserID == 0 {
		response.SendError(c, "未指定用户", nil)
		return
	}
	list, err := userRoleService.FindUserPermissionByUserId(userRole.UserID)

	// var list []model.UserRole
	// err := db.Sql.Model(&model.Role{}).Where("role_id = ?", userRole.UserID).Find(&list).Error

	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &list)
}
