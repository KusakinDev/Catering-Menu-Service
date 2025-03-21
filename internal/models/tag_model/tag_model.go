package tagmodel

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

type Tag struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	TagDish string `gorm:"type:varchar(50)"`
}

func (tag *Tag) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Tag{})
	if err != nil {
		log.Println("Error migrate Tag model :")
		return err
	}
	log.Println("Success migrate Tag model :")
	return nil
}
