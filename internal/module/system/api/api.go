package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/logger"
	"server/internal/middleware"
)

type SystemApi struct {
	logger        logger.Logger
	jwtMiddleware *middleware.JwtMiddleware
	userApi       *UserApi
	authApi       *AuthApi
}

func NewSystemApi(
	logger logger.Logger,
	jwtMiddleware *middleware.JwtMiddleware,
	userApi *UserApi,
	authApi *AuthApi,
) *SystemApi {
	return &SystemApi{
		logger:        logger,
		jwtMiddleware: jwtMiddleware,
		userApi:       userApi,
		authApi:       authApi,
	}
}

func (r *SystemApi) InitSystemApi(router *gin.RouterGroup) {
	{
		authRouter := router.Group("auth")
		r.authApi.InitAuthApi(authRouter)
	}

	privateRouter := router.Group("")
	privateRouter.Use(r.jwtMiddleware.Handler())

	{
		userRouter := router.Group("user")
		r.userApi.InitUserApi(userRouter)
	}
}
