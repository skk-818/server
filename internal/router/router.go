package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/module/system/api"
)

type Router struct {
	systemApi *api.SystemApi
	engine    *gin.Engine
}

func NewRouter(systemApi *api.SystemApi) *Router {

	engine := gin.Default()

	systemRouter := engine.Group("system")
	systemApi.InitSystemApi(systemRouter)

	router := &Router{engine: engine, systemApi: systemApi}
	return router
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}
