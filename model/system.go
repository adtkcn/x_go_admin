package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	UserID    int    `gorm:"primaryKey;autoIncrement;not null" json:"user_id"`
	UserName  string `gorm:"not null;size:50;uniqueIndex" json:"user_name" form:"user_name"`
	Password  string `gorm:"not null;size:256" json:"password"  form:"password"`
	Avatar    string `gorm:"default:'';size:200" json:"avatar" form:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 角色
type Role struct {
	RoleID   int    `gorm:"primaryKey;autoIncrement;not null" json:"role_id"`
	RoleName string `gorm:"not null;size:50;comment:角色名称" json:"role_name"`
}

// 用户的角色
type UserRole struct {
	UserID int `gorm:"primaryKey;not null" json:"user_id"`
	RoleID int `gorm:"primaryKey;not null" json:"role_id"`
}

// 权限
type Permission struct {
	PermissionID   int    `gorm:"primaryKey;autoIncrement;not null" json:"permission_id"`
	PermissionName string `gorm:"not null;size:50;comment:权限名称" json:"permission_name"`
	PermissionKey  string `gorm:"not null;uniqueIndex;size:50;comment:权限唯一标识" json:"permission_key"`
}

// 角色的权限
type RolePermission struct {
	RoleID       int `gorm:"primaryKey;not null;comment:角色id" json:"role_id"`
	PermissionID int `gorm:"primaryKey;not null;comment:权限id" json:"permission_id"`
}

// 菜单
type Menu struct {
	MenuID   int    `gorm:"primaryKey;autoIncrement;not null" json:"menu_id"`
	MenuName string `gorm:"not null;size:50;comment:菜单名称" json:"menu_name"`
	ParentID int    `gorm:"not null" json:"parent_id"`
	URL      string `gorm:"not null;size:500" json:"url"`
	Icon     string `gorm:"default:null;size:20" json:"icon"`
}

// 角色的菜单
type RoleMenu struct {
	RoleID int `gorm:"primaryKey;not null;comment:角色id" json:"role_id"`
	MenuID int `gorm:"primaryKey;not null;comment:菜单id" json:"menu_id"`
}
