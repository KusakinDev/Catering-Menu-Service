package nutritionfactmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NutritionFact struct {
	Id            int     `gorm:"primaryKey;autoIncrement"`
	Calories      float32 `gorm:"type:real"`
	Proteins      float32 `gorm:"type:real"`
	Fats          float32 `gorm:"type:real"`
	Carbohydrates float32 `gorm:"type:real"`
}

func (nutFact *NutritionFact) DecodeFromContext(c *gin.Context) error {
	if err := c.ShouldBindJSON(&nutFact); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

func (nutFact *NutritionFact) AddToTable() int {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Create(&nutFact).Error
	if err != nil {
		db.CloseDB()
		logrus.Error("Error add to table: ", err)
		return 503
	}

	db.CloseDB()
	return 200
}

func (nutFact *NutritionFact) UpdateInTable() int {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Save(&nutFact).Error
	if err != nil {
		db.CloseDB()
		logrus.Error("Error update in table: ", err)
		return 503
	}

	db.CloseDB()
	return 200
}

func (nutFuct *NutritionFact) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&NutritionFact{})
	if err != nil {
		logrus.Errorln("Error migrate NutritionFact model :")
		return err
	}
	logrus.Infoln("Success migrate NutritionFact model :")
	return nil
}
