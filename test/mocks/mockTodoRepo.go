package mocks

import (
	"time"

	"app/GoSample/db/entities"

	"gorm.io/gorm"
)

type MockTodoRepo struct{}

func (e MockTodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockTodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	return true
}

func (e MockTodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockTodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	return []entities.Todo{
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Now(), IsCompleted: false },
	}, true
}

type MockCreateTodoFailTodoRepo struct{}

func (e MockCreateTodoFailTodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	return false
}

func (e MockCreateTodoFailTodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	return true
}

func (e MockCreateTodoFailTodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockCreateTodoFailTodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	return []entities.Todo{
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Now(), IsCompleted: false },
	}, true
}

