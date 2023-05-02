package system

import (
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type PermissionController struct{}

// 获取用户的权限
func (p *PermissionController) GetUserPermission(c *gin.Context) {
	var permissions []model.Permission
	model.DB.Find(&permissions)

	response.Send(c, "ok", &gin.H{
		"useProTable": []string{"add", "batchAdd", "export", "batchDelete", "status"},
		"authButton":  []string{"add", "edit", "delete", "import", "export"},
	})
}

// 权限相关操作
func (p *PermissionController) List(c *gin.Context) {
	var permissions []model.Permission
	model.DB.Find(&permissions)

	response.Send(c, "ok", permissions)
}

func (p *PermissionController) Create(c *gin.Context) {
	var permission model.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {

		response.SendError(c, err.Error(), nil)
		return
	}
	model.DB.Create(&permission)
	response.Send(c, "ok", permission)
}

func (p *PermissionController) Update(c *gin.Context) {
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

func (p *PermissionController) Delete(c *gin.Context) {
	id := c.Param("id")
	var permission model.Permission
	if err := model.DB.Where("permission_id=?", id).First(&permission).Error; err != nil {
		response.SendError(c, "not found", nil)
		return
	}
	model.DB.Delete(&permission)
	response.Send(c, "ok", permission)
}
