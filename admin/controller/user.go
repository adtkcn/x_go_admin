package controller

import (
	"net/http"
	"x-gin-admin/model"
	"x-gin-admin/utils/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	model.User
}

func (u *User) addUser(c *gin.Context) {
	// 解析请求体中的表单数据
	var user = User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将密码加密存储到数据库中
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate password hash"})
		return
	}

	// 创建用户实体
	user.Password = string(passwordHash)

	// 将用户实体存储到数据库中
	result := model.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add user"})
		return
	}

	token, err := jwt.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	// 在响应中返回新增的用户
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// 用户登录
func (u *User) login(c *gin.Context) {
	// 解析请求体中的表单数据
	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从数据库中获取用户实体
	var user model.User
	result := model.DB.First(&user, "user_name = ?", form.Username)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// 登录成功，返回登录信息和 token
	token, err := jwt.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}

// 用户相关操作
func (u *User) GetUsers(c *gin.Context) {
	var users []User
	model.DB.Find(&users)
	c.JSON(200, users)
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
	user_id := c.Param("user_id")
	var user User
	if err := model.DB.Where("user_id=?", user_id).First(&user).Error; err != nil {
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
		userRouter.GET("GetUsers", user.GetUsers)
		userRouter.GET("addUser", user.addUser)
		userRouter.GET("login", user.login)

	}
}
