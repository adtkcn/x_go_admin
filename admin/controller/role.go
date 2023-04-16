package controller

import (
	"x-gin-admin/model"

	"github.com/gin-gonic/gin"
)

type Role struct {
	model.Role
}

func (u *Role) GetRoles(c *gin.Context) {
	var Roles []Role
	model.DB.Find(&Roles)
	c.JSON(200, Roles)
}

func (u *Role) CreateRole(c *gin.Context) {
	var Role Role
	if err := c.ShouldBindJSON(&Role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Create(&Role)
	c.JSON(201, Role)
}

func (u *Role) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var Role Role
	if err := model.DB.Where("Role_id=?", id).First(&Role).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&Role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Save(&Role)
	c.JSON(200, Role)
}

func (u *Role) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	var Role Role
	if err := model.DB.Where("Role_id=?", id).First(&Role).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	model.DB.Delete(&Role)
	c.Status(204)
}

func RegisterRoleRoutes(r *gin.RouterGroup) {
	role := Role{}
	// 子级路由组
	roleRouter := r.Group("/Role")
	{
		roleRouter.GET("/", role.GetRoles)

	}
}
