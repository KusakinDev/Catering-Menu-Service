package dishmodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Dish struct {
	Id                int    `gorm:"primaryKey;autoIncrement"`
	Type_id           int    `gorm:"not null"`
	Category_id       int    `gorm:"not null"`
	Nutrition_fact_id int    `gorm:"not null"`
	Tag_id            int    `gorm:"not null"`
	Recipe            string `gorm:"type:varchar(100)"`
}

func (dish *Dish) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Dish{})
	if err != nil {
		log.Println("Error migrate Dish model :")
		return err
	}
	log.Println("Success migrate Dish model :")
	return nil
}
