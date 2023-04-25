package handler

import (
	"fmt"
	"log"
	"x-gin-admin/admin/service/userService"
	"x-gin-admin/utils/jwt"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

func GetUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 example 变量
		var user, err = jwt.ParseToken(c.Query("token"))
		if err != nil {
			log.Println(err)
			c.Abort()
			c.JSON(401, gin.H{"code": "401", "message": "用户未登录"})
			return
		}
		if user.UserID == 0 {
			c.Abort()
			response.SendError(c, "未登录", nil)
			return
		}
		c.Set("UserID", user.UserID)
		// 请求前
		c.Next()
	}
}
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

		UserID := c.GetInt("UserID")
		if UserID == 0 {
			c.Abort()
			response.SendError(c, "未登录", nil)
			return
		}
		user := userService.GetUserInfo(UserID)
		if user == nil {
			c.Abort()
			response.SendError(c, "获取用户失败", nil)
			return
		}
		c.Set("UserInfo", user)
		fmt.Println("UserInfo", user)
		// 请求前

		c.Next()
	}
}
