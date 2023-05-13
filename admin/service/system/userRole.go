package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
)

type UserRoleService struct{}

// 获取用户的角色
func (u *UserRoleService) FindByUserId(UserID int) (list *[]model.UserRole, err error) {
	err = db.Sql.Model(&model.UserRole{}).Where("user_id = ?", UserID).Find(&list).Error

	if err != nil {
		return list, err
	}

	return list, err
}

type result struct {
	UserID       int
	RoleID       int
	PermissionID int
}

// 获取用户的角色和权限
func (u *UserRoleService) FindUserPermissionByUserId(UserID int) (res result, err error) {
	// var userRole *[]model.UserRole

	// var res result
	err = db.Sql.Model(&model.UserRole{}).Joins("left join x_role_permissions on x_role_permissions.role_id = x_user_roles.role_id").Where("user_id = ?", UserID).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, err
}
