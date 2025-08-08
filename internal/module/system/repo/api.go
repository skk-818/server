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

func (r *apiRepo) Create(ctx context.Context, api *model.Api) error {
	err := r.db.WithContext(ctx).Create(api).Error
	return errors.WithStack(err)
}

func (r *apiRepo) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&model.Api{}, id).Error
	return errors.WithStack(err)
}

func (r *apiRepo) Update(ctx context.Context, api *model.Api) error {
	err := r.db.WithContext(ctx).Updates(api).Error
	return errors.WithStack(err)
}

func (r *apiRepo) Find(ctx context.Context, id int64) (*model.Api, error) {
	var api model.Api
	err := r.db.WithContext(ctx).First(&api, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return &api, nil
}

func (r *apiRepo) List(ctx context.Context, req *request.ApiListReq) ([]*model.Api, int64, error) {
	var (
		apis  []*model.Api
		total int64
	)

	db := r.db.WithContext(ctx).Model(&model.Api{})

	if req.Path != "" {
		db = db.Where(model.ApiCol.Path+" LIKE ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		db = db.Where(model.ApiCol.Method+" = ?", req.Method)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	offset, limit := req.BuilderOffsetAndLimit()
	if err := db.
		Offset(offset).
		Limit(limit).
		Order(model.ApiCol.CreatedAt + " DESC").
		Find(&apis).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return apis, total, nil
}

func (r *apiRepo) BatchDelete(ctx context.Context, ids []int64) error {
	err := r.db.WithContext(ctx).Where(model.ApiCol.ID+" IN ?", ids).Delete(&model.Api{}).Error
	return errors.WithStack(err)
}

func (r *apiRepo) FindByIds(ctx context.Context, ids []int64) ([]*model.Api, error) {
	var apis []*model.Api
	err := r.db.WithContext(ctx).Where(model.ApiCol.ID+" IN ?", ids).Find(&apis).Error
	return apis, errors.WithStack(err)
}

func (r *apiRepo) BatchCreate(ctx context.Context, list []*model.Api) error {
	err := r.db.WithContext(ctx).Create(&list).Error
	return errors.WithStack(err)
}
