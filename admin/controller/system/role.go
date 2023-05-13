package system

import (
	"x-gin-admin/admin/service/system"
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

var roleService = system.RoleService{}

func (u *RoleController) List(c *gin.Context) {
	var params model.BaseQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	where := map[string]interface{}{}
	list, err := roleService.FindByPage(params, where)
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "获取成功", list)
}

func (u *RoleController) FindOne(c *gin.Context) {
	var id = c.Query("role_id")
	if id == "" {
		response.SendError(c, "未传入role_id", nil)
		return
	}
	// var list model.Role
	// err := db.Sql.Model(&model.Role{}).Where("role_id = ?", id).First(&list).Error

	role, err := roleService.FindOne(id)
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "获取成功", &role)
}

func (u *RoleController) Create(c *gin.Context) {
	var Role model.Role
	if err := c.ShouldBind(&Role); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	err := db.Sql.Create(&Role).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	response.Send(c, "ok", &Role)
}

func (u *RoleController) Update(c *gin.Context) {
	id := c.PostForm("role_id")
	var Role model.Role
	if err := db.Sql.Where("role_id=?", id).First(&Role).Error; err != nil {
		response.SendError(c, "未找到角色", nil)
		return
	}
	if err := c.ShouldBind(&Role); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	err := db.Sql.Save(&Role).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", &Role)
}

func (u *RoleController) Delete(c *gin.Context) {
	id := c.PostForm("role_id")

	var Role model.Role
	if err := db.Sql.Where("role_id=?", id).First(&Role).Error; err != nil {
		response.SendError(c, "没找到角色", nil)
		return
	}
	err := roleService.DeleteById(id)
	// err := db.Sql.Delete(&Role).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "删除成功", &id)
}
