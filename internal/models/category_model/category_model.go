package categorymodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Category struct {
	Id            int    `gorm:"primaryKey;autoIncrement"`
	Category_dish string `gorm:"type:varchar(20)"`
}

func (cat *Category) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Category{})
	if err != nil {
		log.Println("Error migrate Category model :")
		return err
	}
	log.Println("Success migrate Category model :")
	return nil
}
