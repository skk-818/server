package api

import "github.com/gin-gonic/gin"

type AuthApi struct{}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

func (a *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", a.Login)
}

func (a *AuthApi) Login(ctx *gin.Context) {

}
