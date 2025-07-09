package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/module/system/usecase"
)

type RoleApi struct {
	logger      logger.Logger
	roleUsecase *usecase.RoleUsecase
}

func NewRoleApi(logger logger.Logger, roleUsecase *usecase.RoleUsecase) *RoleApi {
	return &RoleApi{
		logger:      logger,
		roleUsecase: roleUsecase,
	}
}

func (a *RoleApi) InitRoleApi(router *gin.RouterGroup) {

}
