package api

import (
	"github.com/gin-gonic/gin"
	"server/pkg/response"
)

type AuthApi struct{}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

func (a *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", a.Login)
}

func (a *AuthApi) Login(c *gin.Context) {
	response.Success(c)
}
