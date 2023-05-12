package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

type QueryParams struct {
	Key      string `form:"key"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

func (u *RoleController) List(c *gin.Context) {
	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	query := db.Sql.Model(&model.Role{})
	if params.Key != "" {
		query = query.Where("role_name LIKE ?", "%"+params.Key+"%")
	}

	var list []model.Role
	var count int64
	offset := (params.Page - 1) * params.PageSize
	query.Offset(offset).Limit(params.PageSize).Count(&count).Find(&list)
	response.Send(c, "ok", &response.QueryRes{
		Count: count,
		List:  &list,
		Page:  params.Page,
	})
}

func (u *RoleController) Create(c *gin.Context) {
	var Role model.Role
	if err := c.ShouldBind(&Role); err != nil {
		// c.JSON(400, gin.H{"error": err.Error()})
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
		c.JSON(404, gin.H{"error": "Role not found"})
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
	err := db.Sql.Delete(&Role).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "删除成功", &id)
}
