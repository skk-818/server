package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
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

func (r *userRepo) Find(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Preload(model.UserCol.Roles).
		First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	return errors.WithStack(err)
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&model.User{}, id).Error
	return errors.WithStack(err)
}

func (r *userRepo) List(ctx context.Context, req *request.UserListReq) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	db := r.db.WithContext(ctx).Model(&model.User{}).Preload("Roles")

	if req.Username != nil {
		db = db.Where("username LIKE ?", "%"+*req.Username+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Gender != nil {
		db = db.Where("gender = ?", *req.Gender)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	err = db.Order("id DESC").
		Limit(req.PageSize).
		Offset(offset).
		Find(&users).Error

	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return users, total, nil
}

func (r *userRepo) BatchDelete(ctx context.Context, ids []int64) error {
	err := r.db.WithContext(ctx).
		Where("id IN (?)", ids).
		Delete(&model.User{}).Error
	return errors.WithStack(err)
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

func (r *userRepo) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Preload(model.UserCol.Roles).
		Where(model.UserCol.Phone+" = ?", phone).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}
