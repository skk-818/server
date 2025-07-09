package repo

import (
	"context"
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

func (a apiRepo) Create(ctx context.Context, api *model.Api) error {
	//TODO implement me
	panic("implement me")
}

func (a apiRepo) Delete(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (a apiRepo) Update(ctx context.Context, api *model.Api) error {
	//TODO implement me
	panic("implement me")
}

func (a apiRepo) Find(ctx context.Context, i int64) (*model.Api, error) {
	//TODO implement me
	panic("implement me")
}

func (a apiRepo) List(ctx context.Context, req *request.ApiListReq) ([]*model.Api, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a apiRepo) BatchDelete(ctx context.Context, int64s []int64) error {
	//TODO implement me
	panic("implement me")
}
