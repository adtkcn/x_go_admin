package system

import (
	"x-gin-admin/model"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (u *RoleController) List(c *gin.Context) {
	var Roles []model.Role
	model.DB.Find(&Roles)
	c.JSON(200, Roles)
}

func (u *RoleController) Create(c *gin.Context) {
	var Role model.Role
	if err := c.ShouldBindJSON(&Role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Create(&Role)
	c.JSON(200, Role)
}

func (u *RoleController) Update(c *gin.Context) {
	id := c.Param("id")
	var Role model.Role
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

func (u *RoleController) Delete(c *gin.Context) {
	id := c.Param("id")
	var Role model.Role
	if err := model.DB.Where("Role_id=?", id).First(&Role).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	model.DB.Delete(&Role)
	c.Status(204)
}
