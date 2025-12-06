package api

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz"
	"server/internal/module/system/model"
	_ "server/internal/module/system/model/reply"
	_ "server/internal/module/system/model/response"
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

// GetMenuTree godoc
// @Summary 获取菜单树
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} server_internal_module_system_model_response.MenuTreeResp
// @Router /api/system/menu/tree [get]
func (a *MenuApi) GetMenuTree(c *gin.Context) {
	tree, err := a.menuUsecase.GetMenuTree(c)
	if err != nil {
		a.logger.Error("[MenuApi] GetMenuTree error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, tree)
}

// List godoc
// @Summary 获取菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} server_internal_module_system_model_reply.ListMenuReply
// @Router /api/system/menu/list [get]
func (a *MenuApi) List(c *gin.Context) {
	menus, err := a.menuUsecase.List(c, nil)
	if err != nil {
		a.logger.Error("[MenuApi] List error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, menus)
}

// Create godoc
// @Summary 创建菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body model.Menu true "菜单信息"
// @Success 200 {string} string "success"
// @Router /api/system/menu [post]
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

// Update godoc
// @Summary 更新菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body model.Menu true "菜单信息"
// @Success 200 {string} string "success"
// @Router /api/system/menu [put]
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

// Delete godoc
// @Summary 删除菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "菜单ID"
// @Success 200 {string} string "success"
// @Router /api/system/menu/{id} [delete]
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
