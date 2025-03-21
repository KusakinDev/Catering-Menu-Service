package nutritionfactmodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Nutrition_fact struct {
	Id            int
	Calories      float32
	Proteins      float32
	Fats          float32
	Carbohydrates float32
}

func (nutFuct *Nutrition_fact) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Nutrition_fact{})
	if err != nil {
		log.Println("Error migrate Nutrition_fact model :")
		return err
	}
	log.Println("Success migrate Nutrition_fact model :")
	return nil
}
