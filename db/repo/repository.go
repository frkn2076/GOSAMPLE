package repo

import (
	"app/GoSample/db"
	"gorm.io/gorm"
)

func BeginTransaction() *gorm.DB {
	tx := db.GormDB.Begin()
	return tx
}

func CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
