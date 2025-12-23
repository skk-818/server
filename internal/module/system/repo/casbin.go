package repo

import (
	"server/internal/core/mysql"
	"server/internal/module/system/biz/repo"

	"gorm.io/gorm"
)

type casbinRepo struct {
	db *gorm.DB
}

func NewCasbinRepo(systemDB *mysql.SystemDB) repo.CasbinRepo {
	return &casbinRepo{
		db: systemDB.DB,
	}
}

func (cr *casbinRepo) AdapterDB() *gorm.DB {
	return cr.db
}
