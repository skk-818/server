package repo

import (
	"context"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/usecase/repo"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) repo.RoleRepo {
	return &roleRepo{db: db}
}

func (r *roleRepo) Create(ctx context.Context, role *model.Role) error {
	return nil
}
