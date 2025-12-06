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

type RoleApi struct {
	logger      logger.Logger
	roleUsecase *biz.RoleUsecase
}

func NewRoleApi(logger logger.Logger, roleUsecase *biz.RoleUsecase) *RoleApi {
	return &RoleApi{
		logger:      logger,
		roleUsecase: roleUsecase,
	}
}

func (a *RoleApi) InitRoleApi(router *gin.RouterGroup) {
	router.GET("list", a.List)
	router.POST("", a.Create)
	router.PUT("", a.Update)
	router.DELETE(":id", a.Delete)
}

// List godoc
// @Summary 获取角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param current query int false "页码"
// @Param size query int false "每页数量"
// @Param name query string false "角色名称"
// @Success 200 {object} server_internal_module_system_model_reply.ListRoleReply
// @Router /api/system/role/list [get]
func (a *RoleApi) List(c *gin.Context) {
	var req request.RoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, err)
		return
	}
	result, err := a.roleUsecase.List(c, &req)
	if err != nil {
		a.logger.Error("[RoleApi] List error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, result)
}

// Create godoc
// @Summary 创建角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.CreateRoleReq true "角色信息"
// @Success 200 {string} string "success"
// @Router /api/system/role [post]
func (a *RoleApi) Create(c *gin.Context) {
	var req request.CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.roleUsecase.Create(c, &req); err != nil {
		a.logger.Error("[RoleApi] Create error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// Update godoc
// @Summary 更新角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.UpdateRoleReq true "角色信息"
// @Success 200 {string} string "success"
// @Router /api/system/role [put]
func (a *RoleApi) Update(c *gin.Context) {
	var req request.UpdateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.roleUsecase.Update(c, &req); err != nil {
		a.logger.Error("[RoleApi] Update error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// Delete godoc
// @Summary 删除角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {string} string "success"
// @Router /api/system/role/{id} [delete]
func (a *RoleApi) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.roleUsecase.Delete(c, idInt); err != nil {
		a.logger.Error("[RoleApi] Delete error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}
