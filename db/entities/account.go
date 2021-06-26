package entities

import "gorm.io/gorm"

type Account struct {
	Id       uint   `gorm:"uniqueIndex;autoIncrement:true"`
	UserName string `gorm:"primaryKey"`
	Password string
	gorm.Model
}

func (Account) TableName() string {
	return "Account"
}
