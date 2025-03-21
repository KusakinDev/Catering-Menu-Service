package typemodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Type struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	TypeDish string `gorm:"type:varchar(20)"`
}

func (typ *Type) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Type{})
	if err != nil {
		log.Println("Error migrate Type model :")
		return err
	}
	log.Println("Success migrate Type model :")
	return nil
}
