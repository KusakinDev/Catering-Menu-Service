package allergytypemodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type AllergyType struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Type string `gorm:"type:varchar(20)"`
}

func (all_type *AllergyType) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&AllergyType{})
	if err != nil {
		log.Println("Error migrate Allergy_type model :")
		return err
	}
	log.Println("Success migrate Allergy_type model :")
	return nil
}
