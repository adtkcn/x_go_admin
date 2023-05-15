package system

import (
	"x-gin-admin/db"
	"x-gin-admin/model"
)

type RolePermissionService struct{}

//	type RolePermission struct{
//		model.Permission
//		model.
//	}
//
// 获取角色的权限
func (u *RolePermissionService) FindByRoleIds(RoleIds []int) (list []model.Permission, Ids []int, err error) {
	err = db.Sql.Model(&model.Permission{}).Joins("left join x_role_permissions on x_permissions.permission_id = x_role_permissions.permission_id").Where("x_role_permissions.role_id IN ?", RoleIds).Find(&list).Error

	if err != nil {
		return list, Ids, err
	}

	for i := 0; i < len(list); i++ {
		Ids = append(Ids, list[i].PermissionID)
	}
	return list, Ids, err
}
