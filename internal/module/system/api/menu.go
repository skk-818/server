package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/module/system/usecase"
)

type MenuApi struct {
	logger      logger.Logger
	menuUsecase *usecase.MenuUsecase
}

func NewMenuApi(logger logger.Logger, menuUsecase *usecase.MenuUsecase) *MenuApi {
	return &MenuApi{
		logger:      logger,
		menuUsecase: menuUsecase,
	}
}

func (a *MenuApi) InitMenuApi(router *gin.RouterGroup) {

}
