package allergytypemodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type AllergyType struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Type string `gorm:"type:varchar(50)"`
}

func (all_type *AllergyType) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&AllergyType{})
	if err != nil {
		logrus.Error("Error migrate AllergyType model :")
		return err
	}
	logrus.Info("Success migrate AllergyType model :")
	return nil
}
