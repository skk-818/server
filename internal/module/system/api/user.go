package api

import (
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	"server/pkg"
	"server/pkg/errorx"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
	config      *config.Config
	logger      logger.Logger
	userUsecase *biz.UserUsecase
}

func NewUserApi(config *config.Config, logger logger.Logger, userUsecase *biz.UserUsecase) *UserApi {
	return &UserApi{
		config:      config,
		logger:      logger,
		userUsecase: userUsecase,
	}
}

func (a *UserApi) InitUserApi(router *gin.RouterGroup) {
	router.GET("info", a.Info)
}

func (a *UserApi) Info(c *gin.Context) {
	userId := pkg.GetUserID(c)
	if userId == 0 {
		response.Fail(c, errorx.ErrUnauthorized)
		return
	}

	userInfo, err := a.userUsecase.GetInfo(c, int(userId))
	if err != nil {
		a.logger.Error("[UserApi] GetUserInfo error", zap.Any("userId", userId), zap.Error(err))
		response.Fail(c, err)
		return
	}

	response.SuccessWithData(c, userInfo)
}
