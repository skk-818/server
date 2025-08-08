package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase"
	"server/pkg/response"
)

type ApiApi struct {
	logger     logger.Logger
	apiUsecase *usecase.ApiUsecase
}

func NewApiApi(logger logger.Logger, apiUsecase *usecase.ApiUsecase) *ApiApi {
	return &ApiApi{
		logger:     logger,
		apiUsecase: apiUsecase,
	}
}

func (a *ApiApi) InitApiApi(router *gin.RouterGroup) {
	router.POST("list", a.List)
	router.POST("detail", a.Detail)
	router.POST("create", a.Create)
	router.POST("delete", a.Delete)
	router.POST("update", a.Update)
}

func (a *ApiApi) List(c *gin.Context) {
	var req request.ApiListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[ApiApi.List] bind json error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	reply, err := a.apiUsecase.List(c, &req)
	if err != nil {
		a.logger.Error("[ApiApi.List] apiUsecase.List error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	response.SuccessWithData(c, reply)
}

func (a *ApiApi) Detail(c *gin.Context) {
	var req request.ApiDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[ApiApi.Detail]  bind json error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
	}

	detail, err := a.apiUsecase.Detail(c, &req)
	if err != nil {
		a.logger.Error("[ApiApi.Detail] apiUsecase.Detail error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	response.SuccessWithData(c, detail)
}

func (a *ApiApi) Update(c *gin.Context) {
	var req request.UpdateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[ApiApi.Update]  bind json error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	if err := a.apiUsecase.Update(c, &req); err != nil {
		a.logger.Error("[ApiApi.Update] apiUsecase.Update error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	response.Success(c)
}

func (a *ApiApi) Delete(c *gin.Context) {
	var req request.DeleteApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[ApiApi.Delete] bind json error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	if err := a.apiUsecase.Delete(c, &req); err != nil {
		a.logger.Error("[ApiApi.Delete]  apiUsecase.Delete error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	response.Success(c)
}

func (a *ApiApi) Create(c *gin.Context) {
	var req request.CreateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("[ApiApi.Create] bind json error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}

	if err := a.apiUsecase.Create(c, &req); err != nil {
		a.logger.Error("[ApiApi.Create]  apiUsecase.Create error:", zap.Any("req", req), zap.Any("err", err))
		response.Fail(c, err)
		return
	}
}
