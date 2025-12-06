package api

import (
	"server/internal/core/config"
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	_ "server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/pkg"
	"server/pkg/errorx"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
	config      *config.Config
	logger      logger.Logger
	userUsecase *biz.UserUsecase
}

func NewUserApi(config *config.Config, logger logger.Logger, userUsecase *biz.UserUsecase) *UserApi {
	return &UserApi{
		config:      config,
		logger:      logger,
		userUsecase: userUsecase,
	}
}

func (a *UserApi) InitUserApi(router *gin.RouterGroup) {
	router.GET("info", a.Info)
	router.GET("list", a.List)
	router.POST("", a.Create)
	router.DELETE("", a.Delete)
}

// Info godoc
// @Summary 获取当前用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} server_internal_module_system_model_reply.GetUserInfoReply
// @Router /api/system/user/info [get]
func (a *UserApi) Info(c *gin.Context) {
	userId := pkg.GetUserID(c)
	if userId == 0 {
		response.Fail(c, errorx.ErrUnauthorized)
		return
	}

	userInfo, err := a.userUsecase.GetInfo(c, int(userId))
	if err != nil {
		a.logger.Error("[UserApi] GetUserInfo error", zap.Any("userId", userId), zap.Error(err))
		response.Fail(c, err)
		return
	}

	response.SuccessWithData(c, userInfo)
}

// List godoc
// @Summary 获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param current query int false "页码" default(1)
// @Param size query int false "每页数量" default(20)
// @Param username query string false "用户名"
// @Param status query int false "状态"
// @Success 200 {object} server_internal_module_system_model_reply.UserListReply
// @Router /api/system/user/list [get]
func (a *UserApi) List(c *gin.Context) {
	var req request.UserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, err)
		return
	}
	result, err := a.userUsecase.List(c, &req)
	if err != nil {
		a.logger.Error("[UserApi] List error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, result)
}

// Create godoc
// @Summary 创建用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.CreateUserReq true "用户信息"
// @Success 200 {string} string "success"
// @Router /api/system/user [post]
func (a *UserApi) Create(c *gin.Context) {
	var req request.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.userUsecase.Create(c, &req); err != nil {
		a.logger.Error("[UserApi] Create error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// Delete godoc
// @Summary 删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.DeleteUserReq true "用户ID列表"
// @Success 200 {string} string "success"
// @Router /api/system/user [delete]
func (a *UserApi) Delete(c *gin.Context) {
	var req request.DeleteUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.userUsecase.Delete(c, &req); err != nil {
		a.logger.Error("[UserApi] Delete error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}
