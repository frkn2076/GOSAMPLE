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
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: false },
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

type MockGetAllFailTodoRepo struct{}

func (e MockGetAllFailTodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockGetAllFailTodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	return true
}

func (e MockGetAllFailTodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockGetAllFailTodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	return nil, false
}

type MockUpdateFailTodoRepo struct{}

func (e MockUpdateFailTodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockUpdateFailTodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	return true
}

func (e MockUpdateFailTodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	return false
}

func (e MockUpdateFailTodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	return []entities.Todo{
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: false },
	}, true
}


type MockDeleteFailTodoRepo struct{}

func (e MockDeleteFailTodoRepo) Create(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockDeleteFailTodoRepo) Delete(db *gorm.DB, id uint, userId uint) bool {
	return false
}

func (e MockDeleteFailTodoRepo) Update(db *gorm.DB, todo entities.Todo) bool {
	return true
}

func (e MockDeleteFailTodoRepo) GetAll(userId string) ([]entities.Todo, bool) {
	return []entities.Todo{
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: false },
	}, true
}