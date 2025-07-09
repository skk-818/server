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
	roleApi          *RoleApi
	apiApi           *ApiApi
	menuApi          *MenuApi
}

func NewSystemApi(
	jwtMiddleware *middleware.JwtMiddleware,
	casbinMiddle *middleware.CasbinMiddleware,
	userApi *UserApi,
	authApi *AuthApi,
	roleApi *RoleApi,
	apiApi *ApiApi,
	menuApi *MenuApi,
) *SystemApi {
	return &SystemApi{
		jwtMiddleware:    jwtMiddleware,
		casbinMiddleware: casbinMiddle,
		userApi:          userApi,
		authApi:          authApi,
		roleApi:          roleApi,
		apiApi:           apiApi,
		menuApi:          menuApi,
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

	{
		roleRouter := router.Group("role")
		r.roleApi.InitRoleApi(roleRouter)
	}

	{
		apiRouter := router.Group("api")
		r.apiApi.InitApiApi(apiRouter)
	}

	{
		menuRouter := router.Group("menu")
		r.menuApi.InitMenuApi(menuRouter)
	}
}
