package allergymodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Allergy struct {
	Id              int `gorm:"primaryKey;autoIncrement"`
	Dish_id         int `gorm:"not null"`
	Allergy_type_id int `gorm:"not null"`
}

func (all *Allergy) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Allergy{})
	if err != nil {
		log.Println("Error migrate Allergy model :")
		return err
	}
	log.Println("Success migrate Allergy model :")
	return nil
}
