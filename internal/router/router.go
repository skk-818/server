package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/module/system/api"
)

func NewRouter(api *api.SystemApi) *gin.Engine {

	engine := gin.Default()

	return engine
}
