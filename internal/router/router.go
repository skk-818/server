package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/config"
	"server/internal/middleware"
	commonApi "server/internal/module/common/api"
	systemApi "server/internal/module/system/api"
)

type Router struct {
	engine    *gin.Engine
	cfg       *config.HTTPServer
	cors      *middleware.CorsMiddleware
	systemApi *systemApi.SystemApi
	commonApi *commonApi.CommonApi
}

func NewRouter(cfg *config.HTTPServer, corsMiddleware *middleware.CorsMiddleware, systemApi *systemApi.SystemApi, commonApi *commonApi.CommonApi) *Router {
	engine := gin.Default()

	if cfg.Cors.Enabled {
		engine.Use(corsMiddleware.Handler())
	}

	systemRouter := engine.Group("api/system")
	systemApi.InitSystemApi(systemRouter)

	commonRouter := engine.Group("api/common")
	commonApi.InitCommonApi(commonRouter)

	router := &Router{engine: engine, systemApi: systemApi}
	return router
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}
