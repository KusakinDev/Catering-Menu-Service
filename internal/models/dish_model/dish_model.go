package dishmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	"github.com/sirupsen/logrus"
)

type Dish struct {
	Id              int    `gorm:"primaryKey;autoIncrement"`
	Name            string `gorm:"type:varchar(50)"`
	TypeId          int    `gorm:"not null"`
	CategoryId      int    `gorm:"not null"`
	NutritionFactId int    `gorm:"not null"`
	Tag_id          int    `gorm:"not null"`
	Recipe          string `gorm:"type:varchar(100)"`
}

func (dish *Dish) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Dish{})
	if err != nil {
		logrus.Errorln("Error migrate Dish model :")
		return err
	}
	logrus.Infoln("Success migrate Dish model :")
	return nil
}
