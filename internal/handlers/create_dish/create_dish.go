package createdish

import (
	"log"

	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	nutritionfactmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/nutrition_fact_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateDish(c *gin.Context) (int, string) {

	var dish dishmodel.Dish
	var nutFact nutritionfactmodel.NutritionFact
	var err error
	var code int

	err = dish.DecodeFromContext(c)
	if err != nil {
		return 400, "Error decode JSON"
	}

	nutFact = dish.NutFuct

	log.Println("dish", dish)
	log.Println("nutFact", nutFact)

	code = nutFact.AddToTable()
	if code != 200 {
		logrus.Println("Err add nutFact")
		return code, "Error add to table"
	}

	dish.NutFuct.Id = nutFact.Id

	code = dish.AddToTable()
	if code != 200 {
		logrus.Println("Err add AddToTable")
		return code, "Error add to table"
	}

	return 200, "Success create new dish"
}
