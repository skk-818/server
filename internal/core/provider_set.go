package core

import (
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/core/mysql"
	"server/internal/core/redis"
	"server/internal/core/router"
	"server/internal/core/server"
	"server/internal/module/system/biz"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	server.NewHTTPServer,
	NewInitManagerProvider,

	router.NewRouter,
	wire.Bind(new(server.EngineProvider), new(*router.Router)),

	config.LoadConfig,
	config.ProvideLoggerConfig,
	config.ProvideHttpServerConfig,
	config.ProviderCorsConfig,
	config.ProvideJwtConfig,
	config.ProvideRedisConfig,

	NewSystemDBProvider,
	NewImDBProvider,
	redis.NewRedis,

	logger.NewZapLogger,
	wire.Bind(new(logger.Logger), new(*logger.ZapLogger)),
)

// NewSystemDBProvider 提供系统数据库连接
func NewSystemDBProvider(cfg *config.Config) (*mysql.SystemDB, error) {
	return mysql.NewSystemDB(cfg.SystemMySQL)
}

// NewImDBProvider 提供IM数据库连接
func NewImDBProvider(cfg *config.Config) (*mysql.ImDB, error) {
	return mysql.NewImDB(cfg.ImMySQL)
}

// NewInitManagerProvider 初始化管理器
func NewInitManagerProvider(router *router.Router, initUsecase *biz.InitUsecase, cronUsecase *biz.CronUsecase) []server.InitManager {
	return []server.InitManager{
		router,
		initUsecase,
		cronUsecase,
	}
}
