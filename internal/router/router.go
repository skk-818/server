package router

import (
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/middleware"
	im "server/internal/module/im/api"
	system "server/internal/module/system/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
)

type Group struct {
	logger    logger.Logger
	cfg       *config.HTTPServer
	cors      *middleware.CorsMiddleware
	systemApi *system.SystemApi
	imApi     *im.IMApi
}

func NewGroup(
	logger logger.Logger,
	cfg *config.HTTPServer,
	corsMiddleware *middleware.CorsMiddleware,
	systemApi *system.SystemApi,
	imApi *im.IMApi,
) *Group {
	return &Group{
		logger:    logger,
		cfg:       cfg,
		cors:      corsMiddleware,
		systemApi: systemApi,
		imApi:     imApi,
	}
}

func (a *Group) InitRouter(engine *gin.Engine) error {
	if a.cfg.Cors.Enabled {
		engine.Use(a.cors.Handler())
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		systemRouter := engine.Group("api/system")
		a.systemApi.InitSystemApi(systemRouter)
	}

	return nil
}
