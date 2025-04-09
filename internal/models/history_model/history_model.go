package historymodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	"github.com/sirupsen/logrus"
)

type ActionEnum string

const (
	CREATE ActionEnum = "CREATE"
	DELETE ActionEnum = "DELETE"
	UPDATE ActionEnum = "UPDATE"
)

type History struct {
	Id     int        `gorm:"primaryKey;autoIncrement"`
	Date   string     `gorm:"type:varchar(50)"`
	Action ActionEnum `gorm:"type:varchar(50)"`
	DishId int        `gorm:"not null"`

	Dish dishmodel.Dish `gorm:"foreignKey:DishId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (history *History) SetActionCREATE() {
	history.Action = CREATE
}

func (history *History) SetActionDELETE() {
	history.Action = DELETE
}

func (history *History) SetActionUPDATE() {
	history.Action = UPDATE
}

func (history *History) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&history).Error
	if err != nil {
		logrus.Error("Error add to table: ", err)
		return 503
	}

	return 200
}

func (history *History) AddAllToTable(histories []History) int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	for _, his := range histories {
		err := db.Connection.Create(&his).Error
		if err != nil {
			logrus.Error("Error add to table: ", err)
			return 503
		}
	}

	return 200
}

func (history *History) GetFromTableById() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Preload("Dish").
		Preload("Dish.Type").
		Preload("Dish.Category").
		Preload("Dish.NutFuct").
		Preload("Dish.Tag").
		First(&history).Error
	if err != nil {
		return 503
	}

	return 200
}

func (history *History) GetAllFromTable() (int, []History) {
	var db database.DataBase
	db.InitDB()

	var allHistory []History

	err := db.Connection.Preload("Dish").
		Preload("Dish.Type").
		Preload("Dish.Category").
		Preload("Dish.NutFuct").
		Preload("Dish.Tag").
		Find(&allHistory).Error
	if err != nil {
		db.CloseDB()
		return 503, []History{}
	}
	db.CloseDB()
	return 200, allHistory
}

func (history *History) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&History{})
	if err != nil {
		logrus.Errorln("Error migrate History model :")
		return err
	}
	logrus.Infoln("Success migrate History model :")
	return nil
}
