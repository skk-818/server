package repo

import (
	"gorm.io/gorm"
	"server/internal/module/system/usecase/repo"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return userRepo{
		db: db,
	}
}
