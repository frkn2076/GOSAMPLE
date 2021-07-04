package repo

import (
	"app/GoSample/db"
	
	"gorm.io/gorm"
)

var Repo *Repository

func init() {
	Repo = new(Repository)
}

type Repository struct{}

func (u *Repository) BeginTransaction() *gorm.DB {
	tx := db.GormDB.Begin()
	return tx
}

func (u *Repository) CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func (u *Repository) RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
