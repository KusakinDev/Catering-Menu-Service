package updatedish

import (
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	nutritionfactmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/nutrition_fact_model"
	"github.com/gin-gonic/gin"
)

func UpdateDish(c *gin.Context) (int, string) {
	var dish dishmodel.Dish
	var nutFact nutritionfactmodel.NutritionFact
	var err error
	var code int

	err = dish.DecodeFromContext(c)
	if err != nil {
		return 400, "Error decode JSON"
	}

	nutFact = dish.NutFuct
	code = nutFact.UpdateInTable()
	if code != 200 {
		return code, "Error update to table"
	}

	dish.NutFuct.Id = nutFact.Id
	code = dish.UpdateInTable()
	if code != 200 {
		return code, "Error update dish"
	}

	return code, "Success update dish"
}
