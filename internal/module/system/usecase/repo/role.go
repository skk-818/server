package repo

import (
	"context"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
)

type RoleRepo interface {
	Create(context.Context, *model.Role) error
	FindByKey(context.Context, string) (*model.Role, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.Role) error
	FindByID(context.Context, int64) (*model.Role, error)
	List(context.Context, *request.RoleListReq) ([]*model.Role, int64, error)
	BatchDelete(context.Context, []int64) error
	FindByIDs(context.Context, []int64) ([]*model.Role, error)
}
