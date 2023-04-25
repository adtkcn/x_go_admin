package main

import (
	"x-gin-admin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server := router.Group("/server")
	{
		routes.Use(server)
	}
	// router.Static("/", "./public")

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
