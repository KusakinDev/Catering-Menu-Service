package allergymodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	allergytypemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/allergy_type_model"
	"github.com/sirupsen/logrus"
)

type Allergy struct {
	Id            int `gorm:"primaryKey;autoIncrement"`
	DishId        int `gorm:"not null"`
	AllergyTypeId int `gorm:"not null"`

	AllergyType allergytypemodel.AllergyType `gorm:"foreignKey:AllergyTypeId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
