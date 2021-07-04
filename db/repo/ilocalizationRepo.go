package repo


import (
	"app/GoSample/db/entities"
	"gorm.io/gorm"
)

type ILocalizationRepo interface {
	Create(db *gorm.DB, localization entities.Localization)
	First(resource string, language string) (entities.Localization, error)
	Update(db *gorm.DB, resource string, message string, language string)
	Delete(db *gorm.DB, resource string, language string)
	GetAll() ([]entities.Localization, error)
}