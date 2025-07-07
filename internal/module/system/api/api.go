package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/middleware"
)

type SystemApi struct {
	jwtMiddleware    *middleware.JwtMiddleware
	casbinMiddleware *middleware.CasbinMiddleware
	userApi          *UserApi
	authApi          *AuthApi
}

func NewSystemApi(
	jwtMiddleware *middleware.JwtMiddleware,
	casbinMiddle *middleware.CasbinMiddleware,
	userApi *UserApi,
	authApi *AuthApi,
) *SystemApi {
	return &SystemApi{
		jwtMiddleware:    jwtMiddleware,
		casbinMiddleware: casbinMiddle,
		userApi:          userApi,
		authApi:          authApi,
	}
}

func (r *SystemApi) InitSystemApi(router *gin.RouterGroup) {
	{
		authRouter := router.Group("auth")
		r.authApi.InitAuthApi(authRouter)
	}

	privateRouter := router.Group("")
	privateRouter.Use(r.jwtMiddleware.Handler(), r.casbinMiddleware.Handler())

	{
		userRouter := router.Group("user")
		r.userApi.InitUserApi(userRouter)
	}
}
