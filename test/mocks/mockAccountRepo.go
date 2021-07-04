package mocks

import(
	// "app/GoSample/db/repo"
	"app/GoSample/db/entities"
	// "app/GoSample/controllers"

	"gorm.io/gorm"
)

type MockNotRegisteredUserAccountRepo struct{}

func (e MockNotRegisteredUserAccountRepo) IsUserNameExist(userName string) bool {
        return false
}

func (e MockNotRegisteredUserAccountRepo) Create(db *gorm.DB, account entities.Account) uint {
        return 1
}

func (e MockNotRegisteredUserAccountRepo) FirstByUserName(userName string) (entities.Account, bool) {
        return entities.Account{}, false
}

type MockRegisteredUserAccountRepo struct{}

func (e MockRegisteredUserAccountRepo) IsUserNameExist(userName string) bool {
        return true
}

func (e MockRegisteredUserAccountRepo) Create(db *gorm.DB, account entities.Account) uint {
        return 1
}

func (e MockRegisteredUserAccountRepo) FirstByUserName(userName string) (entities.Account, bool) {
        return entities.Account{UserName: "furkan", Password: "12345" }, true
}

type MockAccountCreateFailAccountRepo struct{}

func (e MockAccountCreateFailAccountRepo) IsUserNameExist(userName string) bool {
        return false
}

func (e MockAccountCreateFailAccountRepo) Create(db *gorm.DB, account entities.Account) uint {
        return 0
}

func (e MockAccountCreateFailAccountRepo) FirstByUserName(userName string) (entities.Account, bool) {
        return entities.Account{UserName: "furkan", Password: "12345" }, true
}



