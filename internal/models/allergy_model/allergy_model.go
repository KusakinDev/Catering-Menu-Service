package allergymodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type Allergy struct {
	Id              int `gorm:"primaryKey;autoIncrement"`
	Dish_id         int `gorm:"not null"`
	Allergy_type_id int `gorm:"not null"`
}

func (all *Allergy) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Allergy{})
	if err != nil {
		logrus.Errorln("Error migrate Allergy model :")
		return err
	}
	logrus.Infoln("Success migrate Allergy model :")
	return nil
}
