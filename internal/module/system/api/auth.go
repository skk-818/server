package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/module/system/usecase"
	"server/pkg/response"
)

type AuthApi struct {
	logger      logger.Logger
	authUsecase *usecase.AuthUsecase
}

func NewAuthApi(logger logger.Logger, authUsecase *usecase.AuthUsecase) *AuthApi {
	return &AuthApi{
		logger:      logger,
		authUsecase: authUsecase,
	}
}

func (aa *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", aa.Login)
}

func (aa *AuthApi) Login(c *gin.Context) {
	response.SuccessWithData(c, gin.H{"token": "123456", "refreshToken": "123456"})
}
