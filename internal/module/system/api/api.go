package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/module/system/usecase"
)

type ApiApi struct {
	logger     logger.Logger
	apiUsecase *usecase.ApiUsecase
}

func NewApiApi(logger logger.Logger, apiUsecase *usecase.ApiUsecase) *ApiApi {
	return &ApiApi{
		logger:     logger,
		apiUsecase: apiUsecase,
	}
}

func (a *ApiApi) InitApiApi(router *gin.RouterGroup) {

}
