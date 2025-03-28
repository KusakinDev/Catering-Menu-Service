package nutritionfactmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type NutritionFact struct {
	Id            int     `gorm:"primaryKey;autoIncrement"`
	Calories      float32 `gorm:"type:real"`
	Proteins      float32 `gorm:"type:real"`
	Fats          float32 `gorm:"type:real"`
	Carbohydrates float32 `gorm:"type:real"`
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
