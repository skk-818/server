package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"

	"github.com/gin-gonic/gin"
)

type RoleApi struct {
	logger      logger.Logger
	roleUsecase *biz.RoleUsecase
}

func NewRoleApi(logger logger.Logger, roleUsecase *biz.RoleUsecase) *RoleApi {
	return &RoleApi{
		logger:      logger,
		roleUsecase: roleUsecase,
	}
}

func (a *RoleApi) InitRoleApi(router *gin.RouterGroup) {

}
