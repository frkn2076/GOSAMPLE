package repo

import (
	database "app/GoSample/db"
	"app/GoSample/db/entities"
	"app/GoSample/logger"

	"gorm.io/gorm"
)

var Account *AccountRepo

func init() {
	Account = new(AccountRepo)
}

type AccountRepo struct{}

func (u *AccountRepo) IsUserNameExist(userName string) bool {
	var account entities.Account
	database.GormDB.Find(&account, "user_name = ?", userName)
	return account.Id != 0
}

func (u *AccountRepo) Create(db *gorm.DB, account entities.Account) uint {
	if err := db.Create(&account).Error; err != nil {
		logger.ErrorLog("An error occured while creating account - accountRepo.go - Account:", account, "- Error:", err.Error())
		return 0
	}
	return account.Id
}

func (u *AccountRepo) FirstByUserName(userName string) (entities.Account, bool) {
	var account entities.Account
	if err := database.GormDB.Where("user_name = ?", userName).First(&account).Error; err != nil {
		logger.ErrorLog("An error occured while selecting first by username - accountRepo.go - Account:", account, "- Error:", err.Error())
		return account, false
	}
	return account, true
}
