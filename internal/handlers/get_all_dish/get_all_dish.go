package getalldish

import (
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	"github.com/gin-gonic/gin"
)

func GetAllDish(c *gin.Context) (int, []dishmodel.Dish) {
	var dish dishmodel.Dish

	dishes, code := dish.GetAllFromTable()
	if code != 200 {
		return code, []dishmodel.Dish{}
	}
	return code, dishes
}
