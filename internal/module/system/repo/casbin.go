package repo

import (
	"server/internal/module/system/biz/repo"

	"gorm.io/gorm"
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
