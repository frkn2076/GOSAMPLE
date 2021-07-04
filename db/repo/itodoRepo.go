package repo

import (
	"app/GoSample/db/entities"

	"gorm.io/gorm"
)

type ITodoRepo interface {
	Create(db *gorm.DB, todo entities.Todo) bool
	Delete(db *gorm.DB, id uint, userId uint) bool
	Update(db *gorm.DB, todo entities.Todo) bool
	GetAll(userId string) ([]entities.Todo, bool)
}