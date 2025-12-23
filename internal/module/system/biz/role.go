package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/pkg/errorx"

	"go.uber.org/zap"
)

type RoleUsecase struct {
	logger        logger.Logger
	roleRepo      repo.RoleRepo
	apiRepo       repo.ApiRepo
	roleMenuRepo  repo.RoleMenuRepo
	casbinUsecase casbinUsecase
}

func NewRoleUsecase(
	logger logger.Logger,
	roleRepo repo.RoleRepo,
	apiRepo repo.ApiRepo,
	roleMenuRepo repo.RoleMenuRepo,
	casbinUsecase casbinUsecase,
) *RoleUsecase {
	return &RoleUsecase{
		logger:        logger,
		roleRepo:      roleRepo,
		apiRepo:       apiRepo,
		roleMenuRepo:  roleMenuRepo,
		casbinUsecase: casbinUsecase,
	}
}

func (u *RoleUsecase) AssignApiPermissions(ctx context.Context, roleId int64, apiIds []int64) error {
	role, err := u.roleRepo.FindByID(ctx, roleId)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("roleId", roleId), zap.Error(err))
		return errorx.ErrInternal
	}
	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("roleId", roleId))
		return errorx.ErrRoleNotFound
	}

	apis, err := u.apiRepo.FindByIds(ctx, apiIds)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] apiRepo.FindByIds error", zap.Any("apiIds", apiIds), zap.Error(err))
		return errorx.ErrInternal
	}

	if len(apis) != len(apiIds) {
		u.logger.Error("[ RoleUsecase ] api not found", zap.Any("apiIds", apiIds))
		return errorx.ErrApiNotFound
	}

	var policies [][]string
	for i := range apis {
		policies = append(policies, []string{role.Key, apis[i].Path, apis[i].Method})
	}

	if err = u.casbinUsecase.DeletePermissionsForRole(role.Key); err != nil {
		u.logger.Error("[ RoleUsecase ] casbinUsecase.DeletePermissionsForRole error", zap.Any("roleId", roleId), zap.Error(err))
		return errorx.ErrAddPoliciesFail
	}

	ok, err := u.casbinUsecase.AddPolicies(policies)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] casbinUsecase.AddPolicies error", zap.Any("policies", policies), zap.Error(err))
		return errorx.ErrInternal
	}

	if !ok {
		u.logger.Error("[ RoleUsecase ] casbinUsecase.AddPolicies failed", zap.Any("policies", policies))
		return errorx.ErrAddPoliciesFail
	}

	return nil
}

func (u *RoleUsecase) Create(ctx context.Context, req *request.CreateRoleReq) error {

	role, err := u.roleRepo.FindByKey(ctx, req.Key)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByKey error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}

	if role != nil {
		u.logger.Error("[ RoleUsecase ] role already exists", zap.Any("req", req))
		return errorx.ErrRoleAlreadyExists
	}

	createRole := &model.Role{
		Name:      req.Name,
		Key:       req.Key,
		Status:    model.RoleStatusEnable,
		DataScope: model.RoleDataScopeAll,
		Sort:      *req.Sort,
		IsSystem:  model.RoleNotSystem,
		Remark:    req.Remark,
	}

	if err := u.roleRepo.Create(ctx, createRole); err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.Create error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}

	// TODO 给角色分配访问动态菜单的权限

	return nil
}

func (u *RoleUsecase) Delete(ctx context.Context, id int64) error {
	ids := []int64{id}
	roles, err := u.roleRepo.FindByIDs(ctx, ids)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("id", id), zap.Error(err))
		return errorx.ErrInternal
	}

	if len(roles) != len(ids) {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("id", id))
		return errorx.ErrRoleNotFound
	}

	var deleteIds []int64
	for i := range roles {
		if roles[i].IsSystem != model.RoleIsSystem {
			deleteIds = append(deleteIds, int64(roles[i].ID))
		}
	}

	if len(deleteIds) == 0 {
		u.logger.Error("[ RoleUsecase ] role is system", zap.Any("id", id))
		return errorx.ErrRoleIsSystem
	}

	if err := u.roleRepo.BatchDelete(ctx, deleteIds); err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.BatchDelete error", zap.Any("id", id), zap.Error(err))
		return errorx.ErrInternal
	}

	return nil
}

func (u *RoleUsecase) Update(ctx context.Context, req *request.UpdateRoleReq) error {

	role, err := u.roleRepo.FindByID(ctx, *req.Id)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}

	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("req", req))
		return errorx.ErrRoleNotFound
	}

	if role.IsSystem == model.RoleIsSystem {
		u.logger.Error("[ RoleUsecase ] role is system", zap.Any("req", req))
		return errorx.ErrRoleIsSystem
	}

	role.Name = req.Name
	role.Key = req.Key
	role.Remark = req.Remark
	role.Sort = *req.Sort

	if err := u.roleRepo.Update(ctx, role); err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.Update error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}

	return nil
}

func (u *RoleUsecase) Get(ctx context.Context, req *request.GetRoleReq) (*reply.GetRoleReply, error) {

	role, err := u.roleRepo.FindByID(ctx, *req.Id)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("req", req), zap.Error(err))
		return nil, errorx.ErrInternal
	}

	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("req", req))
		return nil, errorx.ErrRoleNotFound
	}

	return reply.BuilderGetRoleReply(role), nil
}

func (u *RoleUsecase) List(ctx context.Context, req *request.RoleListReq) (*reply.ListRoleReply, error) {

	roles, total, err := u.roleRepo.List(ctx, req)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.List error", zap.Any("req", req), zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return reply.BuilderListRoleReply(roles, total, *req.Page, *req.PageSize), nil
}

// GetRoleApiPermissions 获取角色的API权限列表
func (u *RoleUsecase) GetRoleApiPermissions(ctx context.Context, roleId int64) ([]int64, error) {
	role, err := u.roleRepo.FindByID(ctx, roleId)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("roleId", roleId), zap.Error(err))
		return nil, errorx.ErrInternal
	}
	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("roleId", roleId))
		return nil, errorx.ErrRoleNotFound
	}

	// 从 Casbin 获取角色的所有权限策略
	policies, err := u.casbinUsecase.GetPermissionsForRole(role.Key)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] casbinUsecase.GetPermissionsForRole error", zap.Any("roleId", roleId), zap.Error(err))
		return nil, errorx.ErrInternal
	}

	// 提取所有的 path 和 method，然后查询对应的 API ID
	var pathMethods []struct {
		Path   string
		Method string
	}
	for _, policy := range policies {
		if len(policy) >= 3 {
			pathMethods = append(pathMethods, struct {
				Path   string
				Method string
			}{
				Path:   policy[1],
				Method: policy[2],
			})
		}
	}

	if len(pathMethods) == 0 {
		return []int64{}, nil
	}

	// 根据 path 和 method 查询 API 列表
	apis, err := u.apiRepo.FindByPathMethods(ctx, pathMethods)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] apiRepo.FindByPathMethods error", zap.Any("pathMethods", pathMethods), zap.Error(err))
		return nil, errorx.ErrInternal
	}

	// 提取 API ID
	var apiIds []int64
	for _, api := range apis {
		apiIds = append(apiIds, int64(api.ID))
	}

	return apiIds, nil
}

// AssignMenuPermissions 分配菜单权限给角色
func (u *RoleUsecase) AssignMenuPermissions(ctx context.Context, roleId int64, menuIds []uint64) error {
	role, err := u.roleRepo.FindByID(ctx, roleId)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("roleId", roleId), zap.Error(err))
		return errorx.ErrInternal
	}
	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("roleId", roleId))
		return errorx.ErrRoleNotFound
	}

	// 分配菜单权限
	if err := u.roleMenuRepo.AssignMenus(ctx, uint64(roleId), menuIds); err != nil {
		u.logger.Error("[ RoleUsecase ] roleMenuRepo.AssignMenus error", zap.Any("roleId", roleId), zap.Any("menuIds", menuIds), zap.Error(err))
		return errorx.ErrInternal
	}

	return nil
}

// GetRoleMenuPermissions 获取角色的菜单权限列表
func (u *RoleUsecase) GetRoleMenuPermissions(ctx context.Context, roleId int64) ([]uint64, error) {
	role, err := u.roleRepo.FindByID(ctx, roleId)
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleRepo.FindByID error", zap.Any("roleId", roleId), zap.Error(err))
		return nil, errorx.ErrInternal
	}
	if role == nil {
		u.logger.Error("[ RoleUsecase ] role not found", zap.Any("roleId", roleId))
		return nil, errorx.ErrRoleNotFound
	}

	// 获取菜单权限
	menuIds, err := u.roleMenuRepo.GetMenuIdsByRoleId(ctx, uint64(roleId))
	if err != nil {
		u.logger.Error("[ RoleUsecase ] roleMenuRepo.GetMenuIdsByRoleId error", zap.Any("roleId", roleId), zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return menuIds, nil
}
