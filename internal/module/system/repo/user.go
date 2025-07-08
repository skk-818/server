package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/usecase/repo"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	return errors.WithStack(err)
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Preload(model.UserCol.Roles).
		Where(model.UserCol.Username+" = ?", username).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}
