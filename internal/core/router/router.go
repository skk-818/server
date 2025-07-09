package router

import (
	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		engine *gin.Engine
		Provider
	}

	Provider interface {
		InitRouter(*gin.Engine) error
	}
)

func NewRouter(provider Provider) *Router {
	engine := gin.Default()
	return &Router{engine: engine, Provider: provider}
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}

func (r *Router) Initializer() error {
	return r.Provider.InitRouter(r.engine)
}
