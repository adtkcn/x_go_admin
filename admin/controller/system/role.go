package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (u *RoleController) List(c *gin.Context) {
	var Roles []model.Role
	db.Sql.Find(&Roles)
	c.JSON(200, Roles)
}

func (u *RoleController) Create(c *gin.Context) {
	var Role model.Role
	if err := c.ShouldBindJSON(&Role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Sql.Create(&Role)
	c.JSON(200, Role)
}

func (u *RoleController) Update(c *gin.Context) {
	id := c.Param("id")
	var Role model.Role
	if err := db.Sql.Where("Role_id=?", id).First(&Role).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&Role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Sql.Save(&Role)
	c.JSON(200, Role)
}

func (u *RoleController) Delete(c *gin.Context) {
	id := c.Param("id")
	var Role model.Role
	if err := db.Sql.Where("Role_id=?", id).First(&Role).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	db.Sql.Delete(&Role)
	c.Status(204)
}
