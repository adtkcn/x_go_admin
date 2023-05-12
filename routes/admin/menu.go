package admin

import (
	"x-gin-admin/admin/controller/system"

	"github.com/gin-gonic/gin"
)

func UseMenu(r *gin.RouterGroup) {
	menu := system.MenuController{}
	r.GET("menu/List", menu.List)
	r.POST("menu/Create", menu.Create)
	r.POST("menu/Update", menu.Update)
	r.POST("menu/Delete", menu.Delete)
	r.GET("menu/GetMenusByUser", menu.GetMenusByUser)
}
