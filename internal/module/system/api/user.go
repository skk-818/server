package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/module/system/usecase"
	"server/pkg/response"
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

func (ua *UserApi) InitUserApi(router *gin.RouterGroup) {
	router.POST("info", ua.Info)
}

func (ua *UserApi) Info(c *gin.Context) {
	response.SuccessWithData(c, gin.H{"name": "", "roles": []string{"R_ADMIN"}, "introduction": "", "avatar": "", "email": ""})
}
