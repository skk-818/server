package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/pkg"
	"server/pkg/errorx"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/schema"
)

type InitUsecase struct {
	logger        logger.Logger
	initRepo      repo.InitRepo
	userRepo      repo.UserRepo
	roleRepo      repo.RoleRepo
	menuRepo      repo.MenuRepo
	apiRepo       repo.ApiRepo
	casbinUsecase casbinUsecase
}

func NewInitUsecase(
	logger logger.Logger,
	initRepo repo.InitRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
	menuRepo repo.MenuRepo,
	apiRepo repo.ApiRepo,
	casbinUsecase casbinUsecase,
) *InitUsecase {
	return &InitUsecase{
		logger:        logger,
		initRepo:      initRepo,
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		menuRepo:      menuRepo,
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

	now := time.Now()

	user := &model.User{
		BaseModel:   model.BaseModel{ID: 1},
		Username:    "admin",
		Password:    pkg.HashPassword("123456"),
		Nickname:    "系统管理员",
		Email:       "202000000@qq.com",
		Phone:       "15599999999",
		Gender:      model.UserGenderMale,
		Status:      model.UserStatusEnable,
		IsAdmin:     model.UserIsSystem,
		Province:    "四川省",
		City:        "成都市",
		District:    "xxx",
		Address:     "四川省成都市xxx",
		Position:    "后端开发工程师",
		Department:  "开发部",
		JobTitle:    "开发经理",
		Tags:        strings.Join([]string{"天然呆", "懒癌患者"}, ","),
		Roles:       []*model.Role{role},
		LastLoginAt: &now,
		LastLoginIP: "",
	}
	if err := u.userRepo.Create(context.Background(), user); err != nil {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameUser, "v1.0.0", "初始化超级管理员用户")
}

func (u *InitUsecase) MenuInitialize() error {
	menus := []*model.Menu{
		{BaseModel: model.BaseModel{ID: 1}, ParentID: 0, Name: "Dashboard", Title: "仪表盘", Path: "/dashboard", Component: "/index/index", Roles: model.RoleKeyAdmin, Icon: "ri:pie-chart-line", Sort: 1, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 2}, ParentID: 1, Name: "Console", Title: "工作台", Path: "dashboard/console", Component: "/dashboard/console", Roles: model.RoleKeyAdmin, Icon: "ri:home-smile-2-line", Sort: 1, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 3}, ParentID: 0, Name: "System", Title: "系统管理", Path: "/system", Component: "/index/index", Roles: model.RoleKeyAdmin, Icon: "ri:user-3-line", Sort: 2, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 4}, ParentID: 3, Name: "User", Title: "用户管理", Path: "system/user", Component: "/system/user", Roles: model.RoleKeyAdmin, Icon: "ri:user-line", Sort: 1, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 5}, ParentID: 3, Name: "Role", Title: "角色管理", Path: "system/role", Component: "/system/role", Roles: model.RoleKeyAdmin, Icon: "ri:user-settings-line", Sort: 2, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 6}, ParentID: 3, Name: "Menu", Title: "菜单管理", Path: "system/menu", Component: "/system/menu", Roles: model.RoleKeyAdmin, Icon: "ri:menu-line", Sort: 3, Status: 1, KeepAlive: 1},
		{BaseModel: model.BaseModel{ID: 7}, ParentID: 3, Name: "Api", Title: "接口管理", Path: "system/api", Component: "/system/api", Roles: model.RoleKeyAdmin, Icon: "ri:api-line", Sort: 4, Status: 1, KeepAlive: 1},
	}

	for _, menu := range menus {
		if err := u.menuRepo.Create(context.Background(), menu); err != nil {
			return err
		}
	}

	return u.initRepo.SetInitialized(model.InitNameMenu, "v1.0.0", "初始化菜单")
}

func (u *InitUsecase) ApiInitialize() error {
	apis := []*model.Api{
		{Name: "SystemUserInfo", Path: "/api/system/user/info", Method: "GET", Description: "获取用户信息", Group: "user", Status: 1},
		{Name: "SystemUserList", Path: "/api/system/user/list", Method: "GET", Description: "获取用户列表", Group: "user", Status: 1},
		{Name: "SystemUserCreate", Path: "/api/system/user", Method: "POST", Description: "创建用户", Group: "user", Status: 1},
		{Name: "SystemUserDelete", Path: "/api/system/user", Method: "DELETE", Description: "删除用户", Group: "user", Status: 1},
		{Name: "SystemRoleList", Path: "/api/system/role/list", Method: "GET", Description: "获取角色列表", Group: "role", Status: 1},
		{Name: "SystemRoleCreate", Path: "/api/system/role", Method: "POST", Description: "创建角色", Group: "role", Status: 1},
		{Name: "SystemRoleUpdate", Path: "/api/system/role", Method: "PUT", Description: "更新角色", Group: "role", Status: 1},
		{Name: "SystemRoleDelete", Path: "/api/system/role/*", Method: "DELETE", Description: "删除角色", Group: "role", Status: 1},
		{Name: "SystemMenuTree", Path: "/api/system/menu/tree", Method: "GET", Description: "获取菜单树", Group: "menu", Status: 1},
		{Name: "SystemMenuList", Path: "/api/system/menu/list", Method: "GET", Description: "获取菜单列表", Group: "menu", Status: 1},
		{Name: "SystemMenuCreate", Path: "/api/system/menu", Method: "POST", Description: "创建菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuUpdate", Path: "/api/system/menu", Method: "PUT", Description: "更新菜单", Group: "menu", Status: 1},
		{Name: "SystemMenuDelete", Path: "/api/system/menu/*", Method: "DELETE", Description: "删除菜单", Group: "menu", Status: 1},
		{Name: "SystemApiList", Path: "/api/system/api/list", Method: "GET", Description: "获取API列表", Group: "api", Status: 1},
		{Name: "SystemApiCreate", Path: "/api/system/api", Method: "POST", Description: "创建API", Group: "api", Status: 1},
		{Name: "SystemApiUpdate", Path: "/api/system/api", Method: "PUT", Description: "更新API", Group: "api", Status: 1},
		{Name: "SystemApiDelete", Path: "/api/system/api/*", Method: "DELETE", Description: "删除API", Group: "api", Status: 1},
	}
	if err := u.apiRepo.BatchCreate(context.Background(), apis); err != nil {
		return err
	}
	return u.initRepo.SetInitialized(model.InitNameApi, "v1.0.0", "初始化管理员 api")
}

func (u *InitUsecase) CasbinInitialize() error {
	policies := [][]string{
		{model.RoleKeyAdmin, "/api/system/user/info", "GET"},
		{model.RoleKeyAdmin, "/api/system/user/list", "GET"},
		{model.RoleKeyAdmin, "/api/system/user", "POST"},
		{model.RoleKeyAdmin, "/api/system/user", "DELETE"},
		{model.RoleKeyAdmin, "/api/system/role/list", "GET"},
		{model.RoleKeyAdmin, "/api/system/role", "POST"},
		{model.RoleKeyAdmin, "/api/system/role", "PUT"},
		{model.RoleKeyAdmin, "/api/system/role/*", "DELETE"},
		{model.RoleKeyAdmin, "/api/system/menu/tree", "GET"},
		{model.RoleKeyAdmin, "/api/system/menu/list", "GET"},
		{model.RoleKeyAdmin, "/api/system/menu", "POST"},
		{model.RoleKeyAdmin, "/api/system/menu", "PUT"},
		{model.RoleKeyAdmin, "/api/system/menu/*", "DELETE"},
		{model.RoleKeyAdmin, "/api/system/api/list", "GET"},
		{model.RoleKeyAdmin, "/api/system/api", "POST"},
		{model.RoleKeyAdmin, "/api/system/api", "PUT"},
		{model.RoleKeyAdmin, "/api/system/api/*", "DELETE"},
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
