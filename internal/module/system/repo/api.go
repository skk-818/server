package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
)

type apiRepo struct {
	db *gorm.DB
}

func NewApiRepo(db *gorm.DB) repo.ApiRepo {
	return &apiRepo{db: db}
}

func (a *apiRepo) Create(ctx context.Context, api *model.Api) error {
	err := a.db.WithContext(ctx).Create(api).Error
	return errors.WithStack(err)
}

func (a *apiRepo) Delete(ctx context.Context, id int64) error {
	err := a.db.WithContext(ctx).Delete(&model.Api{}, id).Error
	return errors.WithStack(err)
}

func (a *apiRepo) Update(ctx context.Context, api *model.Api) error {
	err := a.db.WithContext(ctx).Updates(api).Error
	return errors.WithStack(err)
}

func (a *apiRepo) Find(ctx context.Context, id int64) (*model.Api, error) {
	var api model.Api
	err := a.db.WithContext(ctx).First(&api, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &api, nil
}

func (a *apiRepo) List(ctx context.Context, req *request.ApiListReq) ([]*model.Api, int64, error) {
	var (
		apis  []*model.Api
		total int64
	)

	db := a.db.WithContext(ctx).Model(&model.Api{})

	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	if err := db.
		Offset((req.Page - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("created_at DESC").
		Find(&apis).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return apis, total, nil
}

func (a *apiRepo) BatchDelete(ctx context.Context, ids []int64) error {
	err := a.db.WithContext(ctx).Where("id IN ?", ids).Delete(&model.Api{}).Error
	return errors.WithStack(err)
}
