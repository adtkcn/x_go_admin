package system

import (
	"fmt"
	"x-gin-admin/db"
	"x-gin-admin/model"
)

type RoleMenuService struct{}

// 获取角色的菜单
func (u *RoleMenuService) FindByRoleIds(RoleIds []int) (list []model.Menu, Ids []int, err error) {
	fmt.Println("RoleIds", RoleIds)
	err = db.Sql.Debug().Model(&model.Menu{}).Joins("left join x_role_menus on x_menus.menu_id = x_role_menus.menu_id").Where("x_role_menus.role_id IN ?", RoleIds).Find(&list).Error

	if err != nil {
		fmt.Println(err)
		return list, Ids, err
	}

	for i := 0; i < len(list); i++ {
		Ids = append(Ids, list[i].MenuID)
	}
	return list, Ids, err
}
