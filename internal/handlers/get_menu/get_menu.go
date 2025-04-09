package getmenu

import (
	menumodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/menu_model"
	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) (int, []menumodel.Menu) {
	var menu menumodel.Menu
	code, menus := menu.GetAllFromTable()
	if code != 200 {
		return code, []menumodel.Menu{}
	}
	return 200, menus
}
