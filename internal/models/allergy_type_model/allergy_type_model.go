package allergytypemodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Allergy_type struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Allergy_type string `gorm:"type:varchar(20)"`
}

func (all_type *Allergy_type) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Allergy_type{})
	if err != nil {
		log.Println("Error migrate Allergy_type model :")
		return err
	}
	log.Println("Success migrate Allergy_type model :")
	return nil
}
