package mocks

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type MockRepo struct{}

func (e MockRepo) BeginTransaction() *gorm.DB {
	connection := os.Getenv("PGSQLConnection")
	gormDB, _ := gorm.Open(postgres.Open(connection), &gorm.Config{})
	return gormDB
}

func (e MockRepo) CommitTransaction(tx *gorm.DB) {}

func (e MockRepo) RollbackTransaction(tx *gorm.DB) {}
