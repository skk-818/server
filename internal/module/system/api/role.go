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
	router.POST("assign-api-permissions", a.AssignApiPermissions)
	router.GET(":id/api-permissions", a.GetRoleApiPermissions)
	router.POST("assign-menu-permissions", a.AssignMenuPermissions)
	router.GET(":id/menu-permissions", a.GetRoleMenuPermissions)
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

// AssignApiPermissions godoc
// @Summary 分配API权限
// @Description 为角色分配API权限，会先清除该角色的所有权限，然后重新分配
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.AssignApiPermissionsReq true "角色ID和API ID列表"
// @Success 200 {string} string "success"
// @Router /api/system/role/assign-api-permissions [post]
func (a *RoleApi) AssignApiPermissions(c *gin.Context) {
	var req request.AssignApiPermissionsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.roleUsecase.AssignApiPermissions(c, req.RoleId, req.ApiIds); err != nil {
		a.logger.Error("[RoleApi] AssignApiPermissions error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// GetRoleApiPermissions godoc
// @Summary 获取角色的API权限列表
// @Description 获取指定角色已分配的API权限ID列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {array} int64 "API ID列表"
// @Router /api/system/role/{id}/api-permissions [get]
func (a *RoleApi) GetRoleApiPermissions(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, err)
		return
	}
	apiIds, err := a.roleUsecase.GetRoleApiPermissions(c, idInt)
	if err != nil {
		a.logger.Error("[RoleApi] GetRoleApiPermissions error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.SuccessWithData(c, apiIds)
}

// AssignMenuPermissions godoc
// @Summary 分配菜单权限
// @Description 为角色分配菜单权限，会先清除该角色的所有菜单权限，然后重新分配
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body request.AssignMenuPermissionsReq true "角色ID和菜单ID列表"
// @Success 200 {string} string "success"
// @Router /api/system/role/assign-menu-permissions [post]
func (a *RoleApi) AssignMenuPermissions(c *gin.Context) {
	var req request.AssignMenuPermissionsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err)
		return
	}
	if err := a.roleUsecase.AssignMenuPermissions(c, req.RoleId, req.MenuIds); err != nil {
		a.logger.Error("[RoleApi] AssignMenuPermissions error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	response.Success(c)
}

// GetRoleMenuPermissions godoc
// @Summary 获取角色的菜单权限列表
// @Description 获取指定角色已分配的菜单权限ID列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {object} map[string]interface{} "菜单ID列表"
// @Router /api/system/role/{id}/menu-permissions [get]
func (a *RoleApi) GetRoleMenuPermissions(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, err)
		return
	}
	menuIds, err := a.roleUsecase.GetRoleMenuPermissions(c, idInt)
	if err != nil {
		a.logger.Error("[RoleApi] GetRoleMenuPermissions error", zap.Error(err))
		response.Fail(c, err)
		return
	}
	// 返回格式与前端期望一致
	response.SuccessWithData(c, map[string]interface{}{
		"menuIds": menuIds,
	})
}
