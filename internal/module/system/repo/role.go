package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) repo.RoleRepo {
	return &roleRepo{db: db}
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Role{}).Error
	return errors.WithStack(err)
}

func (r *roleRepo) Update(ctx context.Context, role *model.Role) error {
	// 建议使用 Select("*")，更新所有字段（包括零值）
	err := r.db.WithContext(ctx).
		Model(&model.Role{}).
		Where("id = ?", role.ID).
		Select("*").
		Updates(role).Error
	return errors.WithStack(err)
}

func (r *roleRepo) FindByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&role).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &role, nil
}

func (r *roleRepo) List(ctx context.Context, req *request.RoleListReq) ([]*model.Role, int64, error) {
	var (
		roles []*model.Role
		total int64
		db    = r.db.WithContext(ctx).Model(&model.Role{})
	)

	if req.Name != "" {
		db = db.Where(model.RoleCol.Name+" LIKE ?", "%"+req.Name+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	offset, limit := req.BuilderOffsetAndLimit()
	err = db.Order(model.RoleCol.Sort + " ASC").
		Limit(limit).
		Offset(offset).
		Find(&roles).Error

	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return roles, total, nil
}

func (r *roleRepo) BatchDelete(ctx context.Context, ids []int64) error {
	err := r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Delete(&model.Role{}).Error
	return errors.WithStack(err)
}

func (r *roleRepo) Create(ctx context.Context, role *model.Role) error {
	err := r.db.WithContext(ctx).Create(role).Error
	return errors.WithStack(err)
}

func (r *roleRepo) FindByKey(ctx context.Context, key string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Where(model.RoleCol.Key+" = ?", key).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &role, nil
}

func (r *roleRepo) FindByIDs(ctx context.Context, ids []int64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).Where(model.RoleCol.ID+" IN ?", ids).Find(&roles).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return roles, nil
}

func (r *roleRepo) FindByKeys(ctx context.Context, keys []string) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).Where(model.RoleCol.Key+" IN ?", keys).Find(&roles).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return roles, nil
}
