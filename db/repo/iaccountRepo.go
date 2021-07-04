package repo

import (
	"app/GoSample/db/entities"

	"gorm.io/gorm"
)

type IAccountrepo interface {
	IsUserNameExist(userName string) bool
	Create(db *gorm.DB, account entities.Account) uint
	FirstByUserName(userName string) (entities.Account, bool)
}
