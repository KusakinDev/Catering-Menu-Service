package nutritionfactmodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
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
		log.Println("Error migrate NutritionFact model :")
		return err
	}
	log.Println("Success migrate NutritionFact model :")
	return nil
}
