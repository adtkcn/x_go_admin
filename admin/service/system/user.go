package system

import (
	"errors"
	"x-gin-admin/model"
	"x-gin-admin/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// 返回token
func (u *UserService) Login(username, password string) (string, error) {
	var user model.User
	result := model.DB.First(&user, "user_name = ?", username)
	if result.Error != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		// response.SendError(c, "invalid username or password", nil)
		return "", errors.New("invalid username or password")
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return "", errors.New("invalid password or username")
	}

	// 登录成功，返回登录信息和 token
	token, err := jwt.GenerateToken(user.UserID)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return "", err
	}
	return token, nil

}

// 用户相关信息
func (u *UserService) GetUserInfo(UserID int) *model.User {
	if UserID == 0 {
		return nil
	}
	// 可以先从redis获取
	// 失败再从sql读取

	var users = model.User{
		UserID: UserID,
	}
	model.DB.Find(&users)

	return &users
}
