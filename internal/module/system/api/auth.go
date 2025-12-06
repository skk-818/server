package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	_ "server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/pkg/errorx"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct {
	logger      logger.Logger
	authUsecase *biz.AuthUsecase
}

func NewAuthApi(logger logger.Logger, authUsecase *biz.AuthUsecase) *AuthApi {
	return &AuthApi{
		logger:      logger,
		authUsecase: authUsecase,
	}
}

func (a *AuthApi) InitAuthApi(router *gin.RouterGroup) {
	router.POST("login", a.Login)
	router.POST("register", a.Register)
	router.POST("emailLogin", a.EmailLogin)
}

// Login godoc
// @Summary 用户登录
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param body body request.LoginReq true "登录信息"
// @Success 200 {object} server_internal_module_system_model_reply.LoginReply
// @Router /api/auth/login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[AuthApi] ShouldBindJSON error", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, errorx.ErrInvalidParam)
		return
	}

	reply, err := a.authUsecase.Login(c, &req, c.ClientIP())
	if err != nil {
		a.logger.Error("[AuthApi] Login error", zap.Any("req", req), zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, reply)
}

// Register godoc
// @Summary 用户注册
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param body body request.RegisterReq true "注册信息"
// @Success 200 {string} string "success"
// @Router /api/auth/register [post]
func (a *AuthApi) Register(c *gin.Context) {
	var req request.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[AuthApi] ShouldBindJSON error", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, errorx.ErrInvalidParam)
		return
	}

	if err := a.authUsecase.Register(c, &req); err != nil {
		a.logger.Error("[AuthApi] Register error", zap.Any("req", req), zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// EmailLogin godoc
// @Summary 邮箱登录
// @Tags 认证管理
// @Accept json
// @Produce json
// @Param body body request.EmailLoginReq true "邮箱登录信息"
// @Success 200 {object} server_internal_module_system_model_reply.LoginReply
// @Router /api/auth/emailLogin [post]
func (a *AuthApi) EmailLogin(c *gin.Context) {
	var req request.EmailLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errorx.ErrInvalidParam)
		return
	}

	reply, err := a.authUsecase.EmailLogin(c, &req, c.ClientIP())
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, reply)
}
