package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
)

type SystemApi struct {
	logger  logger.Logger
	userApi *UserApi
}

func NewSystemApi(logger logger.Logger, userApi *UserApi) *SystemApi {
	return &SystemApi{
		logger:  logger,
		userApi: userApi,
	}
}

func (r *SystemApi) InitSystemApi(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		r.logger.Info("ping requested")
		c.JSON(200, gin.H{"message": "pong"})
	})
}
