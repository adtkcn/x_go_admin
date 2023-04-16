package model

import (
	"fmt"
	"x-gin-admin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var DBName = config.Options.Database.DBName
	var DBUser = config.Options.Database.DBUser
	var DBPassword = config.Options.Database.DBPassword
	var DBHost = config.Options.Database.DBHost
	// 连接MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBName)

	fmt.Println("dsn", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建表结构
	err = DB.AutoMigrate(&User{}, &Role{}, &UserRole{}, &Permission{}, &RolePermission{}, &Menu{}, &RoleMenu{})
	if err != nil {
		panic(err)
	}
}
