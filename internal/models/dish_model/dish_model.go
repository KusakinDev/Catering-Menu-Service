package dishmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	categorymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/category_model"
	nutritionfactmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/nutrition_fact_model"
	tagmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/tag_model"
	typemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/type_model"
	pb "github.com/KusakinDev/Catering-Menu-Service/internal/services/get_dish/get_dish_serv"
	"github.com/sirupsen/logrus"
)

type Dish struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(50)"`
	TypeId     int    `gorm:"not null"`
	CategoryId int    `gorm:"not null"`
	NutFactId  int    `gorm:"not null"`
	TagId      int    `gorm:"not null"`
	Recipe     string `gorm:"type:varchar(100)"`

	Type     typemodel.Type                   `gorm:"foreignKey:TypeId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category categorymodel.Category           `gorm:"foreignKey:CategoryId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NutFuct  nutritionfactmodel.NutritionFact `gorm:"foreignKey:NutFactId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tag      tagmodel.Tag                     `gorm:"foreignKey:TagId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (dish *Dish) AddToTable() int {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Create(&dish)
	if err != nil {
		db.CloseDB()
		return 503
	}

	db.CloseDB()
	return 200
}

func (dish *Dish) GetFromTableById() int {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&dish)
	if err != nil {
		db.CloseDB()
		return 503
	}

	db.CloseDB()
	return 200
}

func (dish *Dish) GetAllFromTable() ([]Dish, int) {
	var db database.DataBase
	db.InitDB()

	var dishes []Dish

	err := db.Connection.
		Preload("Type").
		Preload("Category").
		Preload("NutFuct").
		Preload("Tag").
		Find(&dishes).Error
	if err != nil {
		db.CloseDB()
		return []Dish{}, 503
	}
	db.CloseDB()
	return dishes, 200
}

func (dish *Dish) GetFromTableByName() int {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Where("name = ?", dish.Name).First(&dish)
	if err != nil {
		db.CloseDB()
		return 503
	}

	db.CloseDB()
	return 200
}

func (dish *Dish) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Dish{})
	if err != nil {
		logrus.Errorln("Error migrate Dish model :")
		return err
	}
	logrus.Infoln("Success migrate Dish model :")
	return nil
}

func (dish *Dish) ToGRPCDish() *pb.Dish {
	return &pb.Dish{
		Id:   int32(dish.Id),
		Name: dish.Name,
		Type: &pb.Type{
			Id:       int32(dish.Type.Id),
			TypeDish: dish.Type.TypeDish,
		},
		Category: &pb.Category{
			Id:           int32(dish.Category.Id),
			CategoryDish: dish.Category.CategoryDish,
		},
		NutFact: &pb.NutritionFact{
			Id:            int32(dish.NutFuct.Id),
			Calories:      dish.NutFuct.Calories,
			Proteins:      dish.NutFuct.Proteins,
			Fats:          dish.NutFuct.Fats,
			Carbohydrates: dish.NutFuct.Carbohydrates,
		},
		Tag: &pb.Tag{
			Id:      int32(dish.Tag.Id),
			TagDish: dish.Tag.TagDish,
		},
		Recipe: dish.Recipe,
	}
}
