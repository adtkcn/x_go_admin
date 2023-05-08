package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 服务配置
type ServerConfigurations struct {
	Port int
}

// sql 配置
type SqlConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
}

// redis配置
type RedisConfigurations struct {
	Addr     string
	Password string
	DB       int
}

type Configurations struct {
	Server ServerConfigurations
	Sql    SqlConfigurations
	Redis  RedisConfigurations
}

var Options Configurations

func init() {
	config := viper.New()

	config.SetConfigName("config")
	config.AddConfigPath("./")
	config.AutomaticEnv()
	config.SetConfigType("yml")

	//查找并读取配置文件
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	// config.SetDefault("database.dbname", "x_go_admin")
	err := config.Unmarshal(&Options)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic("read config.yaml file error")
	}
	fmt.Print(Options)
	// fmt.Printf("reading using model:\n database=%s,port=%d \n",
	// 	Options.Database.DBName,
	// 	Options.Server.Port,
	// )
	// fmt.Printf("reading without model:\n database=%s,port=%d,path=%s,var=%s \n",
	// 	config.GetString("database.dbname"),
	// 	config.GetInt("server.port"),
	// 	config.GetString("EXAMPLE_PATH"),
	// 	config.GetString("EXAMPLE_VAR"))
}
