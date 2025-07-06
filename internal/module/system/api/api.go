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

}
