package main

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	allergymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/allergy_model"
	categorymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/category_model"
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	tagmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/tag_model"
	typemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/type_model"
)

func main() {

	var db database.DataBase
	db.InitDB()

	var dish dishmodel.Dish
	dish.MigrateToDB(db)

	var all allergymodel.Allergy
	all.MigrateToDB(db)

	var cat categorymodel.Category
	cat.MigrateToDB(db)

	var nutFuct dishmodel.Dish
	nutFuct.MigrateToDB(db)

	var tag tagmodel.Tag
	tag.MigrateToDB(db)

	var typ typemodel.Type
	typ.MigrateToDB(db)

}
