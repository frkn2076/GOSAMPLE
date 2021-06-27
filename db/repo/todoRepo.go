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

func (u *TodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	if err := db.Create(&todo).Error; err != nil {
		logger.ErrorLog("An error occured while creating todo - todoRepo.go - Todo:", todo, "- Error:", err.Error())
		return false
	}
	return true
}

func (u *TodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	if err := db.Where("id = ? and user_id = ?", id, userId).Delete(&entities.Todo{}).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while deleting todo - todoRepo.go - id = %d - Error: %s", id, err.Error())
		logger.ErrorLog(errorMessage)
		return false
	}
	return true
}

func (u *TodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	update := map[string]interface{}{
		"name": todo.Name,
		"description": todo.Description,
		"deadline": todo.Deadline,
		"is_completed": todo.IsCompleted,
	}

	if err := db.Model(&entities.Todo{}).Where("id = ? and user_id = ?", todo.Id, todo.UserId).Updates(update).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while updating todo - todoRepo.go - Error: %s", err.Error())
		logger.ErrorLog(errorMessage)
		return false
	}
	return true
}

func (u *TodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	var todos []entities.Todo
	if err := database.GormDB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		errorMessage := fmt.Sprintf("An error occured while getting all todos - todoRepo.go - Error: %s", err.Error())
		logger.ErrorLog(errorMessage)
		return todos, false
	}
	return todos, true
}
