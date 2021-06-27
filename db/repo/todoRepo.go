package repo

import (
	"fmt"

	database "app/GoSample/db"
	"app/GoSample/db/entities"
	"app/GoSample/logger"

	"gorm.io/gorm"
)

var Todo *TodoRepo

func init() {
	Todo = new(TodoRepo)
}

type TodoRepo struct{}

func (u *TodoRepo) Create(db *gorm.DB, todo entities.Todo) {
	if err := db.Create(&todo).Error; err != nil {
		logger.ErrorLog("An error occured while creating todo - todoRepo.go - Todo:", todo, "- Error:", err.Error())
	}
}

func (u *TodoRepo) Delete(db *gorm.DB, id int) {
	if err := db.Where("id = ?", id).Delete(&entities.Todo{}).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while deleting todo - todoRepo.go - id = %d - Error: %s", id, err.Error())
		logger.ErrorLog(errorMessage)
	}
}

func (u *TodoRepo) UpdateIsCompleted(db *gorm.DB, id int, isCompleted bool) {
	if err := db.Model(&entities.Todo{}).Where("id = ?", id).Update("is_completed", isCompleted).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while updating todo - todoRepo.go - id = %d, isCompleted = %t - Error: %s", id, isCompleted, err.Error())
		logger.ErrorLog(errorMessage)
	}
}

func (u *TodoRepo) GetAll(userId string) ([]entities.Todo, error) {
	var todos []entities.Todo
	if err := database.GormDB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while getting all todos - todoRepo.go - Error: %s", err.Error())
		logger.ErrorLog(errorMessage)
		return todos, err
	}
	return todos, nil
}
