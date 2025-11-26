package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	"server/internal/module/system/model"
	"server/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct {
	logger      logger.Logger
	menuUsecase *biz.MenuUsecase
}

func NewMenuApi(logger logger.Logger, menuUsecase *biz.MenuUsecase) *MenuApi {
	return &MenuApi{
		logger:      logger,
		menuUsecase: menuUsecase,
	}
}

func (a *MenuApi) InitMenuApi(router *gin.RouterGroup) {
	router.GET("tree", a.GetMenuTree)
	router.GET("list", a.List)
	router.POST("", a.Create)
	router.PUT("", a.Update)
	router.DELETE(":id", a.Delete)
}

func (a *MenuApi) GetMenuTree(c *gin.Context) {
	tree, err := a.menuUsecase.GetMenuTree(c)
	if err != nil {
		a.logger.Error("[MenuApi] GetMenuTree error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, tree)
}

func (a *MenuApi) List(c *gin.Context) {
	menus, err := a.menuUsecase.List(c, nil)
	if err != nil {
		a.logger.Error("[MenuApi] List error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, menus)
}

func (a *MenuApi) Create(c *gin.Context) {
	var req model.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.menuUsecase.Create(c, &req); err != nil {
		a.logger.Error("[MenuApi] Create error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

func (a *MenuApi) Update(c *gin.Context) {
	var req model.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.menuUsecase.Update(c, &req); err != nil {
		a.logger.Error("[MenuApi] Update error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

func (a *MenuApi) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.menuUsecase.Delete(c, id); err != nil {
		a.logger.Error("[MenuApi] Delete error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}
