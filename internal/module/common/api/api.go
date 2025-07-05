package api

import "github.com/gin-gonic/gin"

type CommonApi struct {
	authApi *AuthApi
}

func NewCommonApi(authApi *AuthApi) *CommonApi {
	return &CommonApi{
		authApi: authApi,
	}
}

func (a *CommonApi) InitCommonApi(router *gin.RouterGroup) {
	authRouter := router.Group("auth")
	a.authApi.InitAuthApi(authRouter)
}
