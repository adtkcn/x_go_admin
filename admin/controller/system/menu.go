package system

import (
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// 菜单相关操作
func (m *MenuController) GetMenus(c *gin.Context) {
	var menus []model.Menu
	model.DB.Find(&menus)

	response.Send(c, "ok", menus)
}

func (m *MenuController) CreateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Create(&menu)
	response.Send(c, "ok", menu)
}

func (m *MenuController) UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	var menu model.Menu
	if err := model.DB.Where("menu_id=?", id).First(&menu).Error; err != nil {
		response.SendError(c, "Menu not found", nil)
		return
	}
	if err := c.ShouldBindJSON(&menu); err != nil {

		response.SendError(c, err.Error(), nil)
		return
	}
	model.DB.Save(&menu)
	response.Send(c, "ok", menu)
}

func (m *MenuController) DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	var menu model.Menu
	if err := model.DB.Where("menu_id=?", id).First(&menu).Error; err != nil {

		response.SendError(c, "Menu not found", nil)
		return
	}
	model.DB.Delete(&menu)
	response.Send(c, "ok", menu)
}
