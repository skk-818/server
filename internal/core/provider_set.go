package core

import (
	"github.com/google/wire"
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/core/mysql"
	"server/internal/core/redis"
	"server/internal/core/router"
	"server/internal/core/server"
	"server/internal/module/system/usecase"
)

var ProviderSet = wire.NewSet(
	server.NewHTTPServer,
	NewInitManagerProvider,

	router.NewRouter,
	wire.Bind(new(server.EngineProvider), new(*router.Router)),

	config.LoadConfig,
	config.ProvideMysqlConfig,
	config.ProvideLoggerConfig,
	config.ProvideHttpServerConfig,
	config.ProviderCorsConfig,
	config.ProvideJwtConfig,
	config.ProvideRedisConfig,

	mysql.NewMySQL,
	redis.NewRedis,

	logger.NewZapLogger,
	wire.Bind(new(logger.Logger), new(*logger.ZapLogger)),
)

// NewInitManagerProvider 初始化管理器
func NewInitManagerProvider(router *router.Router, initUsecase *usecase.InitUsecase, cronUsecase *usecase.CronUsecase) []server.InitManager {
	return []server.InitManager{
		router,
		initUsecase,
		cronUsecase,
	}
}
