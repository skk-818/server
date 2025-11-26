package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"

	"github.com/gin-gonic/gin"
)

type ApiApi struct {
	logger     logger.Logger
	apiUsecase *biz.ApiUsecase
}

func NewApiApi(logger logger.Logger, apiUsecase *biz.ApiUsecase) *ApiApi {
	return &ApiApi{
		logger:     logger,
		apiUsecase: apiUsecase,
	}
}

func (a *ApiApi) InitApiApi(router *gin.RouterGroup) {

}
