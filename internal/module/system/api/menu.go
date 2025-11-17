package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/module/system/biz"
)

type MenuApi struct {
	logger      logger.Logger
	menuUsecase *biz.MenuUsecase
}

func NewMenuApi(logger logger.Logger, menuUsecase *biz.MenuUsecase) *MenuApi {
	return &MenuApi{
		logger:      logger,
		menuUsecase: menuUsecase,
	}
}

func (a *MenuApi) InitMenuApi(router *gin.RouterGroup) {

}
