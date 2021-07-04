package repo

import "gorm.io/gorm"

type IRepo interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
}