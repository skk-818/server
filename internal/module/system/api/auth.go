package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase"
	"server/pkg/response"
	"server/pkg/xerror"
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

func (a *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", a.Login)
	router.POST("register", a.Register)
}

func (a *AuthApi) Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[AuthApi] ShouldBindJSON error", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, xerror.ErrInvalidParam)
		return
	}

	reply, err := a.authUsecase.Login(c, &req)
	if err != nil {
		a.logger.Error("[AuthApi] Login error", zap.Any("req", req), zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, reply)
}

func (a *AuthApi) Register(c *gin.Context) {
	var req request.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[AuthApi] ShouldBindJSON error", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, xerror.ErrInvalidParam)
		return
	}

	if err := a.authUsecase.Register(c, &req); err != nil {
		a.logger.Error("[AuthApi] Register error", zap.Any("req", req), zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}
