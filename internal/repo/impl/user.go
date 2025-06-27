package impl

import "server/internal/repo"

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.Db) repo.UserRepo {
	return userRepo{}
}
