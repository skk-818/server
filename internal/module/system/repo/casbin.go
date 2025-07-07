package repo

import (
	"gorm.io/gorm"
	"server/internal/module/system/usecase/repo"
)

type casbinRepo struct {
	db *gorm.DB
}

func NewCasbinRepo(db *gorm.DB) repo.CasbinRepo {
	return &casbinRepo{
		db: db,
	}
}

func (cr *casbinRepo) AdapterDB() *gorm.DB {
	return cr.db
}
