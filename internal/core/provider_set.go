package core

import (
	"github.com/google/wire"
	"server/internal/core/config"
	"server/internal/core/server"
)

var ProviderSet = wire.NewSet(
	server.NewHTTPServer,
	config.NewConfig,
)
