package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"

	"github.com/gin-gonic/gin"
)

type RoleService struct{}

func (s *RoleService) FindOne(id string) (role *model.Role, err error) {
	err = db.Sql.Model(&model.Role{}).Where("role_id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (s *RoleService) FindByPage(params model.BaseQueryParams, where map[string]interface{}) (role *gin.H, err error) {
	query := db.Sql.Model(&model.Role{}).Where(where)
	if params.Key != "" {
		query = query.Where("role_name LIKE ?", "%"+params.Key+"%")
	}

	var count int64
	var list []model.Role
	offset := (params.Page - 1) * params.PageSize
	err = query.Offset(offset).Limit(params.PageSize).Count(&count).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &gin.H{
		"count": count,
		"list":  &list,
	}, nil
}

func (s *RoleService) DeleteById(id string) (err error) {
	err = db.Sql.Delete(&model.Role{}, id).Error
	return err
}
