package repo

import (
	"fmt"

	database "app/GoSample/db"
	"app/GoSample/db/entities"
	"app/GoSample/logger"

	"gorm.io/gorm"
)

var Localization *LocalizationRepo

func init() {
	Localization = new(LocalizationRepo)
}

type LocalizationRepo struct{}

func (u *LocalizationRepo) Create(db *gorm.DB, localization entities.Localization) {
	if err := db.Create(&localization).Error; err != nil {
		logger.ErrorLog("An error occured while creating localization - localizationRepo.go - Localization:", localization, "- Error:", err.Error())
	}
}

func (u *LocalizationRepo) First(resource string, language string) (entities.Localization, error) {
	var localization entities.Localization
	if err := database.GormDB.Where("resource = ? and language = ?", resource, language).First(&localization).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while getting first localization - localizationRepo.go - resource = %s and language = %s - Error: %s", resource, language, err.Error())
		logger.ErrorLog(errorMessage)
		return localization, err
	}
	return localization, nil
}

func (u *LocalizationRepo) Update(db *gorm.DB, resource string, message string, language string) {
	if err := db.Model(&entities.Localization{}).Where("resource = ? and language = ?", resource, language).Update("message", message).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while updating localization - localizationRepo.go - resource = %s and language = %s, message = %s - Error: %s", resource, language, message, err.Error())
		logger.ErrorLog(errorMessage)
	}
}

func (u *LocalizationRepo) Delete(db *gorm.DB, resource string, language string) {
	if err := db.Where("resource = ? and language = ?", resource, language).Delete(&entities.Localization{}).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while deleting localization - localizationRepo.go - resource = %s and language = %s - Error: %s", resource, language, err.Error())
		logger.ErrorLog(errorMessage)
	}
}

func (u *LocalizationRepo) GetAll() ([]entities.Localization, error) {
	var localizations []entities.Localization
	if err := database.GormDB.Find(&localizations).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while getting all localizations - localizationRepo.go - Error: %s", err.Error())
		logger.ErrorLog(errorMessage)
		return localizations, err
	}
	return localizations, nil
}
