package typemodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type Type struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	TypeDish string `gorm:"type:varchar(50)"`
}

func (typ *Type) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Type{})
	if err != nil {
		logrus.Errorln("Error migrate Type model :")
		return err
	}
	logrus.Infoln("Success migrate Type model :")
	return nil
}
