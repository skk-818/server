package repo

import (
	"context"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"
)

type UserRepo interface {
	Create(context.Context, *model.User) error
	FindByUsername(context.Context, string) (*model.User, error)
	Find(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
	List(context.Context, *request.UserListReq) ([]*model.User, int64, error)
	BatchDelete(context.Context, []int64) error
}
