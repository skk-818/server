package router

import (
	"github.com/google/wire"
	"server/internal/core/server"
)

var ProviderSet = wire.NewSet(
	NewRouter,
	wire.Bind(new(server.EngineProvider), new(*Router)),
)
