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

func (u *AccountRepo) Create(db *gorm.DB, account entities.Account) {
	if err := db.Create(&account).Error; err != nil {
		logger.ErrorLog("An error occured while creating account - accountRepo.go - Account:", account, "- Error:", err.Error())
	}
}

func (u *AccountRepo) IsExist(account entities.Account) bool {
	var accountCheck entities.Account
	if err := database.GormDB.First(&accountCheck, "user_name = ? AND password = ? ", account.UserName, account.Password).Error; err != nil {
		return true
	}
	return false
}
