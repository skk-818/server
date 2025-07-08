package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/module/system/usecase"
	"server/pkg/response"
)

type AuthApi struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

func (aa *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", aa.Login)
}

func (aa *AuthApi) Login(c *gin.Context) {
	response.SuccessWithData(c, gin.H{"token": "123456", "refreshToken": "123456"})
}
