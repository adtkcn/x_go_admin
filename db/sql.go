package db

import (
	"fmt"
	"x-gin-admin/config"
	"x-gin-admin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Sql *gorm.DB

func init() {
	var DBName = config.Options.Sql.DBName
	var DBUser = config.Options.Sql.DBUser
	var DBPassword = config.Options.Sql.DBPassword
	var DBHost = config.Options.Sql.DBHost
	if DBName == "" {
		fmt.Println("未配置 sql.dbname")
		return
	}
	// 连接MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBName)

	// fmt.Println("dsn", dsn)
	var err error
	Sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "x_",
		},
	})

	if err != nil {
		panic(err)
	}

	// 创建表结构
	err = Sql.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserRole{},
		&model.Permission{},
		&model.RolePermission{},
		&model.Menu{},
		&model.RoleMenu{},
		&model.UploadFile{},
	)
	if err != nil {
		panic(err)
	}
}
func Paginate(params model.BaseQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := params.Page
		pageSize := params.PageSize

		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
