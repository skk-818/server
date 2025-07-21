package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/usecase"
	"server/pkg"
	"server/pkg/response"
	"server/pkg/xerror"
)

type UserApi struct {
	config      *config.Config
	logger      logger.Logger
	userUsecase *usecase.UserUsecase
}

func NewUserApi(config *config.Config, logger logger.Logger, userUsecase *usecase.UserUsecase) *UserApi {
	return &UserApi{
		config:      config,
		logger:      logger,
		userUsecase: userUsecase,
	}
}

func (a *UserApi) InitUserApi(router *gin.RouterGroup) {
	router.POST("info", a.Info)
}

func (a *UserApi) Info(c *gin.Context) {
	userId := pkg.GetUserID(c)
	if userId == 0 {
		response.Fail(c, errorx.ErrUnauthorized)
		return
	}

	userInfo, err := a.userUsecase.GetUserInfo(c, int(userId))
	if err != nil {
		a.logger.Error("[UserApi] GetUserInfo error", zap.Any("userId", userId), zap.Error(err))
		response.Fail(c, err)
		return
	}

	response.SuccessWithData(c, reply.ToUserInfoReply(userInfo))
}
