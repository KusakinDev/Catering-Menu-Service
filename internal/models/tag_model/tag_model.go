package tagmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type Tag struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	TagDish string `gorm:"type:varchar(50)"`
}

func (tag *Tag) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Tag{})
	if err != nil {
		logrus.Errorln("Error migrate Tag model :")
		return err
	}
	logrus.Infoln("Success migrate Tag model :")
	return nil
}
