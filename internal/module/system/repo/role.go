package repo

import (
	"context"
	"github.com/pkg/errors"
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
	err := r.db.WithContext(ctx).Create(role).Error
	return errors.WithStack(err)
}

func (r *roleRepo) FindByKey(ctx context.Context, key string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Where("`key` = ?", key).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &role, nil
}
