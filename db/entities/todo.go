package entities

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id          uint `gorm:"uniqueIndex;autoIncrement:true"`
	UserId      uint
	Name        string
	Description string
	Deadline    time.Time
	IsCompleted bool
	gorm.Model
}

func (Todo) TableName() string {
	return "Todo"
}
