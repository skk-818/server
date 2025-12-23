package repo

import (
	"context"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
)

type ApiRepo interface {
	Create(context.Context, *model.Api) error
	Delete(context.Context, int64) error
	Update(context.Context, *model.Api) error
	Find(context.Context, int64) (*model.Api, error)
	List(context.Context, *request.ApiListReq) ([]*model.Api, int64, error)
	BatchDelete(context.Context, []int64) error
	FindByIds(context.Context, []int64) ([]*model.Api, error)
	BatchCreate(context.Context, []*model.Api) error
	FindByPathMethods(context.Context, []struct {
		Path   string
		Method string
	}) ([]*model.Api, error)
}
