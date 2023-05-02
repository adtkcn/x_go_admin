package system

import (
	"sort"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

type Meta struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	IsLink      int    `json:"isLink"`
	IsHide      int    `json:"isHide"`
	IsFull      int    `json:"isFull"`
	IsAffix     int    `json:"isAffix"`
	IsKeepAlive int    `json:"isKeepalive"`
}
type Route struct {
	Path      string  `json:"path"`
	Name      string  `json:"name"`
	Component string  `json:"component"`
	Meta      Meta    `json:"meta"`
	Children  []Route `json:"children"`
	ParentID  int
	MenuID    int
}

// func main() {
//     people := []Person{
//         {"Alice", 25},
//         {"Bob", 30},
//         {"Charlie", 20},
//         {"David", 35},
//     }
//     sort.Sort(ByAge(people))
//     fmt.Println(people)
// }

func buildTree(menus []Route, parentID int) []Route {
	var tree []Route
	for _, menu := range menus {
		if menu.ParentID == parentID {
			children := buildTree(menus, menu.MenuID)
			if len(children) > 0 {
				menu.Children = children
			}
			tree = append(tree, menu)
		}
	}
	return tree
}

// 获取用户的菜单
func (m *MenuController) GetMenusByUser(c *gin.Context) {
	var menus model.Menus
	model.DB.Find(&menus)
	sort.Sort(model.Menus(menus))

	var routes []Route
	for i := 0; i < len(menus); i++ {
		var route = Route{
			Path:      menus[i].Path,
			Name:      menus[i].Name,
			Component: menus[i].Component,
			ParentID:  menus[i].ParentID,
			MenuID:    menus[i].MenuID,
			Meta: Meta{
				Icon:        menus[i].Icon,
				Title:       menus[i].Title,
				IsLink:      menus[i].IsLink,
				IsHide:      menus[i].IsHide,
				IsFull:      menus[i].IsFull,
				IsAffix:     menus[i].IsAffix,
				IsKeepAlive: menus[i].IsKeepAlive,
			},
		}
		routes = append(routes, route)
	}

	response.Send(c, "ok", buildTree(routes, 0))
}

// 菜单相关操作
func (m *MenuController) List(c *gin.Context) {
	var menus []model.Menu
	model.DB.Find(&menus)
	sort.Sort(model.Menus(menus))
	// response.Send(c, "ok", &menus)
	response.Send(c, "ok", buildTree(menus, 0))
}

func (m *MenuController) Create(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	model.DB.Create(&menu)
	response.Send(c, "ok", &menu)
}

func (m *MenuController) Update(c *gin.Context) {
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
	response.Send(c, "ok", &menu)
}

func (m *MenuController) Delete(c *gin.Context) {
	id := c.Param("id")
	var menu model.Menu
	if err := model.DB.Where("menu_id=?", id).First(&menu).Error; err != nil {

		response.SendError(c, "Menu not found", nil)
		return
	}
	model.DB.Delete(&menu)
	response.Send(c, "ok", &menu)
}
