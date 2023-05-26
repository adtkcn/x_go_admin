package system

import (
	"fmt"
	"x-gin-admin/admin/service/system"
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/captcha"
	"x-gin-admin/utils/jwt"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userService = system.UserService{}

type UserController struct{}

// 注册
func (u *UserController) Register(c *gin.Context) {
	// 解析请求体中的表单数据
	var user = model.User{}
	if err := c.ShouldBind(&user); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	// 将密码加密存储到数据库中
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.SendError(c, "failed to generate password hash", nil)
		return
	}

	// 创建用户实体
	user.Password = string(passwordHash)

	// 将用户实体存储到数据库中
	result := db.Sql.Create(&user)
	if result.Error != nil {
		response.SendError(c, result.Error.Error(), nil)
		return
	}

	access_token, err := jwt.GenerateToken(user.UserID)
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	// 在响应中返回新增的用户
	// c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
	response.Send(c, "ok", gin.H{"user": user, "access_token": access_token})
}

// 用户登录
func (u *UserController) Login(c *gin.Context) {
	var verifyID = c.PostForm("verifyID")
	var verifyCode = c.PostForm("verifyCode")

	isPass := captcha.Verify(verifyID, verifyCode)
	if !isPass {
		response.SendError(c, "验证码错误", nil)
		return
	}
	// 解析请求体中的表单数据
	type Form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var form = &Form{}
	if err := c.ShouldBind(&form); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	token, err := userService.Login(form.Username, form.Password)
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", gin.H{"access_token": token})
}

// 用户登录
func (u *UserController) Logout(c *gin.Context) {

	response.Send(c, "已退出登录", nil)
}

// 获取用户信息
func (u *UserController) GetUsers(c *gin.Context) {
	var UserID = c.GetInt("UserID")
	if UserID == 0 {
		response.SendError(c, "失败", nil)
		return
	}
	var users = userService.FindOne(UserID)
	response.Send(c, "ok", users)
}

func (u *UserController) List(c *gin.Context) {
	var params model.BaseQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	where := map[string]interface{}{}
	list, err := userService.FindByPage(params, where)
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "获取成功", list)
}

func (m *UserController) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	// 将密码加密存储到数据库中
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.SendError(c, "failed to generate password hash", nil)
		return
	}
	// 创建用户实体
	user.Password = string(passwordHash)

	err = db.Sql.Create(&user).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", &user)
}
func (u *UserController) Update(c *gin.Context) {
	id := c.PostForm("user_id")
	var user model.User
	if err := db.Sql.Where("user_id=?", id).First(&user).Error; err != nil {
		fmt.Println(err.Error())
		response.SendError(c, "未找到用户", nil)
		return
	}
	if err := c.ShouldBind(&user); err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}

	err := db.Sql.Omit("password").Save(&user).Error
	if err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	response.Send(c, "ok", &user)
}

func (u *UserController) Delete(c *gin.Context) {
	user_id := c.PostForm("user_id")
	var user model.User
	if err := db.Sql.Where("user_id=?", user_id).First(&user).Error; err != nil {
		response.SendError(c, err.Error(), nil)
		return
	}
	db.Sql.Delete(&user)
	response.Send(c, "ok", user)
}
