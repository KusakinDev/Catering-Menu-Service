package createmenu

import (
	"time"

	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	historymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/history_model"
	menumodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/menu_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateMenu(c *gin.Context) (int, string) {
	var history historymodel.History
	var menu menumodel.Menu
	var dish dishmodel.Dish
	var code int

	var menus []menumodel.Menu
	code, menus = menu.DecodeArrayFromContext(c)
	if code != 200 {
		return code, "Error decode JSON"
	}

	var menusOld []menumodel.Menu
	code, menusOld = menu.GetAllFromTable()
	if code != 200 {
		return code, "Error get old menu"
	}

	var historyLog []historymodel.History
	for _, menu = range menusOld {
		code = menu.RemoveFromTable()
		if code != 200 {
			return code, "Error delete old menu"
		}

		history.Date = time.Now().Format("2006-01-02 15:04:05")
		history.DishId = menu.DishId
		history.SetActionDELETE()
		historyLog = append(historyLog, history)
	}

	for _, menu = range menus {
		logrus.Println(menu)
		dish.Id = menu.DishId
		if dish.IsInTable() {
			code = menu.AddToTable()
			if code != 200 {
				logrus.Error("Error create new menu")
				return code, "Error create new menu"
			}

			history.Date = time.Now().Format("2006-01-02 15:04:05")
			history.DishId = menu.DishId
			history.SetActionCREATE()
			logrus.Println("history", history)
			historyLog = append(historyLog, history)
		}
	}
	history.AddAllToTable(historyLog)
	logrus.Println("historyLog", historyLog)

	resp := NotifNewMenu("http://localhost:8081/GetMenu")
	logrus.Info(resp)

	return 200, "Success create new menu"
}
