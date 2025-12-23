package repo

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCasbinRepo,
	NewInitRepo,
	NewUserRepo,
	NewRoleRepo,
	NewApiRepo,
	NewMenuRepo,
	NewRoleMenuRepo,
)
