package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	_ "server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {
	logger     logger.Logger
	apiUsecase *biz.ApiUsecase
}

func NewApiApi(logger logger.Logger, apiUsecase *biz.ApiUsecase) *ApiApi {
	return &ApiApi{
		logger:     logger,
		apiUsecase: apiUsecase,
	}
}

func (a *ApiApi) InitApiApi(router *gin.RouterGroup) {
	router.GET("list", a.List)
	router.POST("", a.Create)
	router.PUT("", a.Update)
	router.DELETE(":id", a.Delete)
}

// List godoc
// @Summary 获取API列表
// @Tags 接口管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param current query int false "页码" default(1)
// @Param size query int false "每页数量" default(20)
// @Param name query string false "接口名称"
// @Param path query string false "接口路径"
// @Param method query string false "请求方法"
// @Success 200 {object} server_internal_module_system_model_reply.ListApiReply
// @Router /api/system/api/list [get]
func (a *ApiApi) List(c *gin.Context) {
	var req request.ApiListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, err)
		return
	}
	result, err := a.apiUsecase.List(c, &req)
	if err != nil {
		a.logger.Error("[ApiApi] List error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, result)
}

// Create godoc
// @Summary 创建API
// @Tags 接口管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.CreateApiReq true "API信息"
// @Success 200 {string} string "success"
// @Router /api/system/api [post]
func (a *ApiApi) Create(c *gin.Context) {
	var req request.CreateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.apiUsecase.Create(c, &req); err != nil {
		a.logger.Error("[ApiApi] Create error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// Update godoc
// @Summary 更新API
// @Tags 接口管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.UpdateApiReq true "API信息"
// @Success 200 {string} string "success"
// @Router /api/system/api [put]
func (a *ApiApi) Update(c *gin.Context) {
	var req request.UpdateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.apiUsecase.Update(c, &req); err != nil {
		a.logger.Error("[ApiApi] Update error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// Delete godoc
// @Summary 删除API
// @Tags 接口管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "API ID"
// @Success 200 {string} string "success"
// @Router /api/system/api/{id} [delete]
func (a *ApiApi) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, err)
		return
	}
	req := &request.DeleteApiReq{ID: idInt}
	if err := a.apiUsecase.Delete(c, req); err != nil {
		a.logger.Error("[ApiApi] Delete error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}
