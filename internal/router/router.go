package router

import (
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/middleware"
	"server/internal/module/system/api"

	"github.com/gin-gonic/gin"
)

type Group struct {
	logger    logger.Logger
	cfg       *config.HTTPServer
	cors      *middleware.CorsMiddleware
	systemApi *api.SystemApi
}

func NewGroup(
	logger logger.Logger,
	cfg *config.HTTPServer,
	corsMiddleware *middleware.CorsMiddleware,
	systemApi *api.SystemApi,
) *Group {
	return &Group{
		logger:    logger,
		cfg:       cfg,
		cors:      corsMiddleware,
		systemApi: systemApi,
	}
}

func (a *Group) InitRouter(engine *gin.Engine) error {
	if a.cfg.Cors.Enabled {
		engine.Use(a.cors.Handler())
	}

	{
		systemRouter := engine.Group("api/system")
		a.systemApi.InitSystemApi(systemRouter)
	}

	return nil
}
