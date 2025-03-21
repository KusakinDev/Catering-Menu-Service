package dishmodel

import (
	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	categorymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/category_model"
	nutritionfactmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/nutrition_fact_model"
	tagmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/tag_model"
	typemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/type_model"
	"github.com/sirupsen/logrus"
)

type Dish struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(50)"`
	TypeId     int    `gorm:"not null"`
	CategoryId int    `gorm:"not null"`
	NutFactId  int    `gorm:"not null"`
	TagId      int    `gorm:"not null"`
	Recipe     string `gorm:"type:varchar(100)"`

	Type     typemodel.Type                   `gorm:"foreignKey:TypeId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category categorymodel.Category           `gorm:"foreignKey:CategoryId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NutFuct  nutritionfactmodel.NutritionFact `gorm:"foreignKey:NutFactId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tag      tagmodel.Tag                     `gorm:"foreignKey:TagId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
