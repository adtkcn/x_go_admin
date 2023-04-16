package controller

import (
	"x-gin-admin/model"

	"github.com/gin-gonic/gin"
)

type User struct {
	model.User
}

// 用户相关操作
func (u *User) GetUsers(c *gin.Context) {
	var users []User
	model.DB.Find(&users)
	c.JSON(200, users)
}

func (u *User) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Create(&user)
	c.JSON(201, user)
}

func (u *User) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := model.DB.Where("user_id=?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Save(&user)
	c.JSON(200, user)
}

func (u *User) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := model.DB.Where("user_id=?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	model.DB.Delete(&user)
	c.Status(204)
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	user := User{}
	// 子级路由组
	userRouter := r.Group("/user")
	{
		userRouter.GET("/", user.GetUsers)
	}
}
