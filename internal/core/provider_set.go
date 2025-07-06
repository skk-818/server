package core

import (
	"github.com/google/wire"
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/core/mysql"
	"server/internal/core/server"
)

var ProviderSet = wire.NewSet(
	server.NewHTTPServer,

	config.LoadConfig,
	config.ProvideMysqlConfig,
	config.ProvideLoggerConfig,
	config.ProvideHttpServerConfig,
	config.ProviderCorsConfig,
	config.ProvideJwtConfig,

	mysql.NewMySQL,

	logger.NewZapLogger,
	wire.Bind(new(logger.Logger), new(*logger.ZapLogger)),
)
