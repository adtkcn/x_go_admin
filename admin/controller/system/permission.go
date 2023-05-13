package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type PermissionController struct{}

// 获取用户的权限
func (p *PermissionController) GetUserPermission(c *gin.Context) {
	var permissions []model.Permission
	db.Sql.Find(&permissions)

	response.Send(c, "ok", &gin.H{
		"useProTable": []string{"add", "batchAdd", "export", "batchDelete", "status"},
		"authButton":  []string{"add", "edit", "delete", "import", "export"},
	})
}

// 权限相关操作
func (p *PermissionController) List(c *gin.Context) {
	var permissions []model.Permission
	db.Sql.Find(&permissions)

	response.Send(c, "ok", permissions)
}

func (p *PermissionController) Create(c *gin.Context) {
	var permission model.Permission
	if err := c.ShouldBind(&permission); err != nil {

		response.SendError(c, err.Error(), nil)
		return
	}
	err := db.Sql.Create(&permission).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", permission)
}

func (p *PermissionController) Update(c *gin.Context) {
	id := c.Param("id")
	var permission model.Permission
	if err := db.Sql.Where("permission_id=?", id).First(&permission).Error; err != nil {

		response.SendError(c, "not found", nil)
		return
	}
	if err := c.ShouldBind(&permission); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	db.Sql.Save(&permission)
	response.Send(c, "ok", permission)
}

func (p *PermissionController) Delete(c *gin.Context) {
	id := c.Param("id")
	var permission model.Permission
	if err := db.Sql.Where("permission_id=?", id).First(&permission).Error; err != nil {
		response.SendError(c, "not found", nil)
		return
	}
	db.Sql.Delete(&permission)
	response.Send(c, "ok", permission)
}
