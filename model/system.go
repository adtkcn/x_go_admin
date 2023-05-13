package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	UserID    int    `gorm:"primaryKey;autoIncrement;not null" json:"user_id" form:"user_id"`
	UserName  string `gorm:"not null;uniqueIndex;size:50;" json:"user_name" form:"user_name"`
	Password  string `gorm:"not null;size:256" json:"password"  form:"password"`
	Avatar    string `gorm:"default:'';size:200" json:"avatar" form:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 角色
type Role struct {
	RoleID   int    `gorm:"primaryKey;autoIncrement;not null" json:"role_id" form:"role_id"`
	RoleName string `gorm:"not null;uniqueIndex;size:50;comment:角色名称" json:"role_name" form:"role_name"`
	// RoleKey   string `gorm:"not null;uniqueIndex;size:50;comment:角色唯一标识" json:"role_key" form:"role_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 用户的角色
type UserRole struct {
	UserID int `gorm:"primaryKey;not null" json:"user_id" form:"user_id"`
	RoleID int `gorm:"primaryKey;not null" json:"role_id" form:"role_id"`
}

// 权限
type Permission struct {
	PermissionID   int    `gorm:"primaryKey;autoIncrement;not null" json:"permission_id" form:"permission_id"`
	PermissionName string `gorm:"not null;size:50;comment:权限名称" json:"permission_name" form:"permission_name"`
	PermissionKey  string `gorm:"not null;uniqueIndex;size:50;comment:权限唯一标识" json:"permission_key" form:"permission_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// 角色的权限
type RolePermission struct {
	RoleID       int `gorm:"primaryKey;not null;comment:角色id" json:"role_id" form:"role_id"`
	PermissionID int `gorm:"primaryKey;not null;comment:权限id" json:"permission_id" form:"permission_id"`
}

// 菜单
type Menu struct {
	MenuID   int `gorm:"primaryKey;autoIncrement;not null" json:"menu_id" form:"menu_id"`
	ParentID int `gorm:"default:0;comment:父级" json:"parent_id" form:"parent_id"`
	Index    int `gorm:"default:0;comment:排序" json:"index" form:"index"`

	Path      string `gorm:"default:'';size:500;comment:路由路径" json:"path" form:"path"`
	Name      string `gorm:"default:'';size:50;comment:路由名称" json:"name" form:"name" `
	Redirect  string `gorm:"default:'';size:500;comment:重定向" json:"redirect" form:"redirect"`
	Component string `gorm:"default:'';size:500;comment:组件路径" json:"component" form:"component"`

	Icon        string `gorm:"default:'';size:20;comment:菜单图标" json:"icon" form:"icon"`
	Title       string `gorm:"default:'';size:50;comment:菜单名称" json:"title" form:"title"`
	IsLink      int    `gorm:"default:0;size:1;comment:是否link" json:"isLink" form:"isLink"`
	IsHide      int    `gorm:"default:0;size:1;comment:是否隐藏" json:"isHide" form:"isHide"`
	IsFull      int    `gorm:"default:0;size:1;comment:是否全屏(示例：数据大屏页面)" json:"isFull" form:"isFull"`
	IsAffix     int    `gorm:"default:0;size:1;comment:是否固定在 tabs nav" json:"isAffix" form:"isAffix"`
	IsKeepAlive int    `gorm:"default:1;size:1;comment:是否缓存" json:"isKeepAlive" form:"isKeepAlive"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// sort.Sort 排序
// type Menus []Menu
// func (a Menus) Len() int           { return len(a) }
// func (a Menus) Less(i, j int) bool { return a[i].Index > a[j].Index }
// func (a Menus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// 角色的菜单
type RoleMenu struct {
	RoleID int `gorm:"primaryKey;not null;comment:角色id" json:"role_id" form:"role_id"`
	MenuID int `gorm:"primaryKey;not null;comment:菜单id" json:"menu_id" form:"menu_id"`
}

// 文件表
type UploadFile struct {
	gorm.Model
	Name       string `gorm:"index;not null;comment:原始文件名" json:"name" form:"name"`
	Path       string `gorm:"not null;comment:文件路径" json:"path" form:"path"`
	Md5        string `gorm:"index;not null;comment:文件md5" json:"md5" form:"md5"` // 相同MD5文件可以做快速上传
	Permission string `gorm:"not null;comment:权限标识" json:"permission" form:"permission"`
}
