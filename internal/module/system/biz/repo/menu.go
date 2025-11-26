package repo

import (
	"context"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
)

type MenuRepo interface {
	Create(context.Context, *model.Menu) error
	Update(context.Context, *model.Menu) error
	Delete(context.Context, int64) error
	Find(context.Context, int64) (*model.Menu, error)
	List(context.Context, *request.MenuListReq) ([]*model.Menu, int64, error)
	BatchDelete(context.Context, []int64) error
	GetAllEnabled(context.Context) ([]*model.Menu, error)
	GetAll(context.Context) ([]*model.Menu, error)
}
