package repo

import (
	"context"
	"server/internal/module/system/model"
)

type UserRepo interface {
	Create(context.Context, *model.User) error
	FindByUsername(context.Context, string) (*model.User, error)
}
