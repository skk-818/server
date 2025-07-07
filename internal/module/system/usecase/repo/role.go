package repo

import (
	"context"
	"server/internal/module/system/model"
)

type RoleRepo interface {
	Create(context.Context, *model.Role) error
}
