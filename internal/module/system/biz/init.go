package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/pkg"
	"server/pkg/errorx"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm/schema"
)

type InitUsecase struct {
	logger        logger.Logger
	initRepo      repo.InitRepo
	userRepo      repo.UserRepo
	roleRepo      repo.RoleRepo
	apiRepo       repo.ApiRepo
	casbinUsecase casbinUsecase
}

func NewInitUsecase(
	logger logger.Logger,
	initRepo repo.InitRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
	apiRepo repo.ApiRepo,
	casbinUsecase casbinUsecase,
) *InitUsecase {
	return &InitUsecase{
		logger:        logger,
		initRepo:      initRepo,
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		apiRepo:       apiRepo,
		casbinUsecase: casbinUsecase,
	}
}

func (u *InitUsecase) InitIfNeeded() error {
	if err := u.initRepo.AutoMigrate([]schema.Tabler{
		&model.Init{}, &model.Role{}, &model.User{}, &model.Menu{}, &model.Api{},
	}); err != nil {
		u.logger.Error("[InitUsecase] failed to initialize database table structure", zap.Any("err", err))
		return err
	}

	initSteps := []struct {
		name    string
		check   func() bool
		execute func() error
	}{
		{"role", u.RoleIsInitialized, u.RoleInitialize},
		{"user", u.UserIsInitialized, u.UserInitialize},
		{"menu", u.MenuIsInitialized, u.MenuInitialize},
		{"api", u.ApiIsInitialized, u.ApiInitialize},
		{"casbin", u.CasbinIsInitialized, u.CasbinInitialize},
	}

	for _, step := range initSteps {
		if !step.check() {
			if err := step.execute(); err != nil {
				u.logger.Error("[InitUsecase] system initialize "+step.name+" fail", zap.Any("err", err))
				return err
			}
		}
		u.logger.Info("[InitUsecase] " + step.name + " initialized")
	}
	return nil
}

func (u *InitUsecase) isInitialized(name string) bool {
	initialized, err := u.initRepo.IsInitialized(name)
	if err != nil {
		u.logger.Error("[InitUsecase] find system "+name+" initialize flag fail", zap.Any("err", err))
		return false
	}
	return initialized
}

func (u *InitUsecase) RoleIsInitialized() bool   { return u.isInitialized(model.InitNameRole) }
func (u *InitUsecase) UserIsInitialized() bool   { return u.isInitialized(model.InitNameUser) }
func (u *InitUsecase) MenuIsInitialized() bool   { return u.isInitialized(model.InitNameMenu) }
func (u *InitUsecase) ApiIsInitialized() bool    { return u.isInitialized(model.InitNameApi) }
func (u *InitUsecase) CasbinIsInitialized() bool { return u.isInitialized(model.InitNameCasbin) }

func (u *InitUsecase) RoleInitialize() error {
	role := &model.Role{
		BaseModel: model.BaseModel{ID: 1},
		Name:      "超级管理员",
		Key:       model.RoleKeyAdmin,
		Status:    model.RoleStatusEnable,
		DataScope: model.RoleDataScopeAll,
		Sort:      1,
		IsSystem:  model.RoleIsSystem,
		Remark:    "系统初始化超级管理员",
	}
	if err := u.roleRepo.Create(context.Background(), role); err != nil {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameRole, "v1.0.0", "初始化超级管理员角色")
}

func (u *InitUsecase) UserInitialize() error {
	role, err := u.roleRepo.FindByKey(context.Background(), model.RoleKeyAdmin)
	if err != nil || role == nil {
		return errorx.ErrAdminRoleNotFound
	}

	user := &model.User{
		BaseModel:  model.BaseModel{ID: 1},
		Username:   "admin",
		Password:   pkg.HashPassword("123456"),
		Nickname:   "系统管理员",
		Email:      "202000000@qq.com",
		Phone:      "15599999999",
		Gender:     model.UserGenderMale,
		Status:     model.UserStatusEnable,
		IsAdmin:    model.UserIsSystem,
		Province:   "四川省",
		City:       "成都市",
		District:   "xxx",
		Address:    "四川省成都市xxx",
		Position:   "后端开发工程师",
		Department: "开发部",
		JobTitle:   "开发经理",
		Tags:       strings.Join([]string{"天然呆", "懒癌患者"}, ","),
		Roles:      []*model.Role{role},
	}
	if err := u.userRepo.Create(context.Background(), user); err != nil {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameUser, "v1.0.0", "初始化超级管理员用户")
}

func (u *InitUsecase) MenuInitialize() error {
	// TODO MENU初始化
	return u.initRepo.SetInitialized(model.InitNameMenu, "v1.0.0", "初始化菜单")
}

func (u *InitUsecase) ApiInitialize() error {
	apis := []*model.Api{
		{Name: "SystemDashboardData", Path: "/api/system/dashboard/data", Method: "POST", Description: "管理员仪表盘数据", Group: "dashboard", Status: 1},
		{Name: "SystemUserCreate", Path: "/api/system/user/create", Method: "POST", Description: "管理员添加用户", Group: "user", Status: 1},
		{Name: "SystemUserList", Path: "/api/system/user/list", Method: "POST", Description: "管理员获取用户分页列表", Group: "user", Status: 1},
		{Name: "SystemUserUpdate", Path: "/api/system/user/update", Method: "POST", Description: "管理员更新用户信息", Group: "user", Status: 1},
		{Name: "SystemUserInfo", Path: "/api/system/user/info", Method: "GET", Description: "管理员获取用户信息", Group: "user", Status: 1},
		{Name: "SystemUserDelete", Path: "/api/system/user/delete", Method: "POST", Description: "管理员删除用户", Group: "user", Status: 1},
		{Name: "SystemUserSwitchStatus", Path: "/api/system/user/switchStatus", Method: "POST", Description: "管理员切换用户状态", Group: "user", Status: 1},
		{Name: "SystemUserSetRole", Path: "/api/system/user/setRole", Method: "POST", Description: "管理员设置用户角色", Group: "user", Status: 1},
		{Name: "SystemRoleCreate", Path: "/api/system/role/create", Method: "POST", Description: "管理员创建角色", Group: "role", Status: 1},
		{Name: "SystemRoleDelete", Path: "/api/system/role/delete", Method: "POST", Description: "管理员更新角色", Group: "role", Status: 1},
		{Name: "SystemRoleUpdate", Path: "/api/system/role/update", Method: "POST", Description: "管理员删除角色", Group: "role", Status: 1},
		{Name: "SystemRoleInfo", Path: "/api/system/role/info", Method: "POST", Description: "管理员查询角色", Group: "role", Status: 1},
		{Name: "SystemRoleList", Path: "/api/system/role/list", Method: "POST", Description: "管理员查询角色列表", Group: "role", Status: 1},
		{Name: "SystemRoleAssignApiPermissions", Path: "/api/system/role/assignApiPermissions", Method: "POST", Description: "管理员分配角色api权限", Group: "role", Status: 1},
		{Name: "SystemRoleAssignMenuPermissions", Path: "/api/system/role/assignMenuPermissions", Method: "POST", Description: "管理员分配角色菜单权限", Group: "role", Status: 1},
		{Name: "SystemMenuCreate", Path: "/api/system/menu/create", Method: "POST", Description: "管理员创建菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuDelete", Path: "/api/system/menu/delete", Method: "POST", Description: "管理员更新菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuUpdate", Path: "/api/system/menu/update", Method: "POST", Description: "管理员删除菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuInfo", Path: "/api/system/menu/info", Method: "POST", Description: "管理员查询菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuList", Path: "/api/system/menu/list", Method: "POST", Description: "管理员查询菜单列表", Group: "menu", Status: 1},
		{Name: "SystemMenuDynamic", Path: "/api/system/menu/dynamic", Method: "POST", Description: "管理员动态菜单", Group: "menu", Status: 1},
	}
	if err := u.apiRepo.BatchCreate(context.Background(), apis); err != nil {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameApi, "v1.0.0", "初始化管理员 api")
}

func (u *InitUsecase) CasbinInitialize() error {
	policies := [][]string{
		{model.RoleKeyAdmin, "/api/system/user/info", "GET"},
	}

	for _, policy := range policies {
		if exist, err := u.casbinUsecase.HasPolicy(policy); err != nil {
			return err
		} else if exist {
			return errorx.ErrPolicyIsExist
		}
	}

	if ok, err := u.casbinUsecase.BatchAddPolicies(policies); err != nil || !ok {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameCasbin, "v1.0.0", "初始化超级管理员权限")
}
