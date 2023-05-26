package system

import (
	"errors"
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// 返回token
func (u *UserService) Login(username, password string) (string, error) {
	var user model.User
	result := db.Sql.First(&user, "user_name = ?", username)
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
func (u *UserService) FindOne(UserID int) *model.User {
	if UserID == 0 {
		return nil
	}
	// 可以先从redis获取
	// 失败再从sql读取

	var users = model.User{
		UserID: UserID,
	}
	db.Sql.Find(&users)

	return &users
}

func (s *UserService) FindByPage(params model.BaseQuery, where map[string]interface{}) (role *gin.H, err error) {
	query := db.Sql.Model(&model.User{}).Where(where)
	if params.Key != "" {
		query = query.Where("user_name LIKE ?", "%"+params.Key+"%")
	}

	var count int64
	var list []model.User
	// offset := (params.Page - 1) * params.PageSize
	err = query.Select("user_id", "user_name", "avatar", "created_at", "updated_at").Count(&count).Scopes(db.Paginate(params)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &gin.H{
		"pageNum":  params.PageNum,
		"pageSize": params.PageSize,
		"count":    count,
		"list":     &list,
	}, nil
}

func (s *UserService) DeleteById(id string) (err error) {
	err = db.Sql.Delete(&model.Role{}, id).Error
	return err
}
