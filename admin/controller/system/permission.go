package system

import (
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type PermissionController struct{}

// 权限相关操作
func (p *PermissionController) GetPermissions(c *gin.Context) {
	var permissions []model.Permission
	model.DB.Find(&permissions)

	response.Send(c, "ok", permissions)
}

func (p *PermissionController) CreatePermission(c *gin.Context) {
	var permission model.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {

		response.SendError(c, err.Error(), nil)
		return
	}
	model.DB.Create(&permission)
	response.Send(c, "ok", permission)
}

func (p *PermissionController) UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	var permission model.Permission
	if err := model.DB.Where("permission_id=?", id).First(&permission).Error; err != nil {

		response.SendError(c, "not found", nil)
		return
	}
	if err := c.ShouldBindJSON(&permission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	model.DB.Save(&permission)
	response.Send(c, "ok", permission)
}

func (p *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	var permission model.Permission
	if err := model.DB.Where("permission_id=?", id).First(&permission).Error; err != nil {
		response.SendError(c, "not found", nil)
		return
	}
	model.DB.Delete(&permission)
	response.Send(c, "ok", permission)
}
