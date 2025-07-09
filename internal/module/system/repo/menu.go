package repo

import (
	"context"
	"gorm.io/gorm"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
)

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) repo.MenuRepo {
	return &menuRepo{db: db}
}

func (m menuRepo) Create(ctx context.Context, menu *model.Menu) error {
	//TODO implement me
	panic("implement me")
}

func (m menuRepo) Update(ctx context.Context, menu *model.Menu) error {
	//TODO implement me
	panic("implement me")
}

func (m menuRepo) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m menuRepo) Find(ctx context.Context, id int64) (*model.Menu, error) {
	//TODO implement me
	panic("implement me")
}

func (m menuRepo) List(ctx context.Context, req *request.MenuListReq) ([]*model.Menu, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m menuRepo) BatchDelete(ctx context.Context, ids []int64) error {
	//TODO implement me
	panic("implement me")
}
