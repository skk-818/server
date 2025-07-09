package router

import (
	"github.com/google/wire"
	"server/internal/core/router"
)

var ProviderSet = wire.NewSet(
	NewGroup,
	wire.Bind(new(router.Provider), new(*Group)),
)
