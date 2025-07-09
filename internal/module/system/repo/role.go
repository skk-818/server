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

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) Update(ctx context.Context, role *model.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) FindByID(ctx context.Context, id int64) (*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) List(ctx context.Context, req *request.RoleListReq) ([]*model.Role, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) BatchDelete(ctx context.Context, ids []int64) error {
	//TODO implement me
	panic("implement me")
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
	err := r.db.WithContext(ctx).Where(model.RoleCol.Key+" = ?", key).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &role, nil
}
