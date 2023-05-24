package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
)

type UserRoleService struct{}

// var rolePermissionService = RolePermissionService{}
var menuService = RoleMenuService{}

// 获取用户的角色
func (u *UserRoleService) FindByUserId(UserID int) (list []model.Role, RoleIds []int, err error) {
	err = db.Sql.Model(&model.Role{}).Joins("left join x_user_roles on x_user_roles.role_id = x_roles.role_id").Where("x_user_roles.user_id = ?", UserID).Find(&list).Error

	if err != nil {
		return list, RoleIds, err
	}

	for i := 0; i < len(list); i++ {
		RoleIds = append(RoleIds, list[i].RoleID)
	}
	return list, RoleIds, err
}

// 获取用户的角色和权限
func (u *UserRoleService) FindUserPermissionByUserId(UserID int) (res map[string]interface{}, err error) {

	Roles, RoleIds, err := u.FindByUserId(UserID)
	if err != nil {
		return res, err
	}
	// permission, permissionIds, _ := rolePermissionService.FindByRoleIds(RoleIds)
	if err != nil {
		return res, err
	}
	menu, menuIds, _ := menuService.FindByRoleIds(RoleIds)

	return map[string]interface{}{
		"Roles":   Roles,
		"RoleIds": RoleIds,
		// "permission":    permission,
		// "permissionIds": permissionIds,
		"menu":    menu,
		"menuIds": menuIds,
	}, err
}
