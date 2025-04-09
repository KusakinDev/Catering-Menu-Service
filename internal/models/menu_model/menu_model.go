package menumodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Menu struct {
	Id     int `gorm:"primaryKey;autoIncrement"`
	DishId int `json:"dish_id" gorm:"not null"`

	Dish dishmodel.Dish `gorm:"foreignKey:DishId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (menu *Menu) DecodeFromContext(c *gin.Context) error {
	if err := c.ShouldBindJSON(&menu); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

func (menu *Menu) DecodeArrayFromContext(c *gin.Context) (int, []Menu) {
	var menus []Menu
	if err := c.ShouldBindJSON(&menus); err != nil {
		logrus.Error("Error decoding JSON array: ", err)
		return 400, []Menu{}
	}
	logrus.Println(menus)
	return 200, menus
}

func (menu *Menu) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&menu).Error
	if err != nil {
		logrus.Error("Error add to table: ", err)
		return 503
	}

	return 200
}

func (menu *Menu) RemoveFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Delete(&menu).Error
	if err != nil {
		logrus.Error("Error delete from table: ", err)
		return 503
	}

	return 200
}

func (menu *Menu) GetFromTableById() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Preload("Dish").
		Preload("Dish.Type").
		Preload("Dish.Category").
		Preload("Dish.NutFuct").
		Preload("Dish.Tag").
		First(&menu).Error
	if err != nil {
		return 503
	}

	return 200
}

func (menu *Menu) GetAllFromTable() (int, []Menu) {
	var db database.DataBase
	db.InitDB()

	var allMenu []Menu

	err := db.Connection.Preload("Dish").
		Preload("Dish.Type").
		Preload("Dish.Category").
		Preload("Dish.NutFuct").
		Preload("Dish.Tag").
		Find(&allMenu).Error
	if err != nil {
		db.CloseDB()
		return 503, []Menu{}
	}
	db.CloseDB()
	return 200, allMenu
}

func (menu *Menu) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Menu{})
	if err != nil {
		logrus.Errorln("Error migrate Menu model :")
		return err
	}
	logrus.Infoln("Success migrate Menu model :")
	return nil
}
