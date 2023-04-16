package main

import (
	"x-gin-admin/admin"
	"x-gin-admin/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server := router.Group("/server")
	{
		// var user = controller.User{}
		admin.RegisterRoutes(server)
		api.RegisterRoutes(server)

	}
	// router.Static("/", "./public")

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
