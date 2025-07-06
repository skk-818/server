package api

import "github.com/gin-gonic/gin"

type CommonApi struct {
	authApi *AuthApi
	userApi *UserApi
}

func NewCommonApi(authApi *AuthApi, userApi *UserApi) *CommonApi {
	return &CommonApi{
		authApi: authApi,
		userApi: userApi,
	}
}

func (a *CommonApi) InitCommonApi(router *gin.RouterGroup) {
	authRouter := router.Group("auth")
	a.authApi.InitAuthApi(authRouter)

	userRouter := router.Group("user")
	a.userApi.InitUserRouter(userRouter)
}
