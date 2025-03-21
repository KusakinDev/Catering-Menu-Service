package categorymodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type Category struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	CategoryDish string `gorm:"type:varchar(50)"`
}

func (cat *Category) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Category{})
	if err != nil {
		logrus.Errorln("Error migrate Category model :")
		return err
	}
	logrus.Infoln("Success migrate Category model :")
	return nil
}
