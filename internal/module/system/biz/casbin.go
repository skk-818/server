package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/pkg/errorx"
	"sync"

	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	CasbinUsecase struct {
		logger   logger.Logger
		enforcer *casbin.Enforcer
		roleRepo repo.RoleRepo
	}

	casbinUsecase interface {
		HasPolicy([]string) (bool, error)
		AddPolicy([]string) (bool, error)
		DeletePermissionsForRole(role string) error
		AddPolicies([][]string) (bool, error)
		BatchAddPolicies([][]string) (bool, error)
	}
)

var (
	once     sync.Once
	enforcer *casbin.Enforcer
)

func NewCasbinUsecase(logger logger.Logger, casbinRepo repo.CasbinRepo, roleRepo repo.RoleRepo) (*CasbinUsecase, error) {
	var err error
	once.Do(func() {
		logger.Info("开始初始化 Casbin Enforcer...")
		db := casbinRepo.AdapterDB()
		adapter, e := gormadapter.NewAdapterByDB(db)
		if e != nil {
			logger.Error("初始化 Casbin Adapter 失败", zap.Any("error", e))
			err = e
			return
		}

		logger.Info("Casbin Adapter 初始化完成")
		modelText := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`
		m, e := casbinModel.NewModelFromString(modelText)
		if e != nil {
			logger.Error("加载 Casbin 模型失败", zap.Any("error", e))
			err = e
			return
		}

		logger.Info("Casbin 模型加载成功")
		enforcer, e = casbin.NewEnforcer(m, adapter)
		if e != nil {
			logger.Error("创建 Casbin Enforcer 失败", zap.Any("error", e))
			err = e
			return
		}

		logger.Info("Casbin Enforcer 创建成功")
		if e := enforcer.LoadPolicy(); e != nil {
			logger.Error("加载 Casbin 策略失败", zap.Any("error", e))
			err = e
			return
		}
		logger.Info("Casbin 策略加载成功")
	})

	if err != nil {
		logger.Error("CasbinUsecase 初始化失败", zap.Any("error", err))
		return nil, err
	}
	logger.Info("CasbinUsecase 初始化完成")
	return &CasbinUsecase{
		logger:   logger,
		enforcer: enforcer,
		roleRepo: roleRepo,
	}, nil
}

// Enforce roleKey path method
func (u *CasbinUsecase) Enforce(sub, obj, act string) (bool, error) {
	role, err := u.roleRepo.FindByKey(context.Background(), model.RoleKeyAdmin)
	if err != nil {
		u.logger.Error("[ CasbinUsecase ] find role by key fail", zap.String("key", sub), zap.Any("error", err))
		return false, err
	}

	if role.Status != model.RoleStatusEnable {
		u.logger.Error("[ CasbinUsecase ] role not enable", zap.Any("role", role))
		return false, errorx.ErrRoleIsDisabled
	}

	return u.enforcer.Enforce(sub, obj, act)
}

func (u *CasbinUsecase) HasPolicy(policy []string) (bool, error) {
	return u.enforcer.HasPolicy(policy)
}

func (u *CasbinUsecase) AddPolicy(policy []string) (bool, error) {
	return u.enforcer.AddPolicy(policy)
}

func (u *CasbinUsecase) BatchAddPolicies(policies [][]string) (bool, error) {
	// 判断 policies 是否为空
	if len(policies) == 0 {
		return false, nil
	}

	// 批量添加策略
	added, err := u.enforcer.AddPolicies(policies)
	if err != nil {
		return false, err
	}

	// 如果 added 为 false，说明有部分策略已经存在，没有全部成功添加
	if !added {
		return false, errors.New("部分策略添加失败，可能已存在")
	}

	// 添加成功后持久化
	err = u.enforcer.SavePolicy()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *CasbinUsecase) DeletePermissionsForRole(role string) error {
	_, err := u.enforcer.RemoveFilteredPolicy(0, role)
	if err != nil {
		u.logger.Error("删除角色权限失败", zap.String("role", role), zap.Error(err))
		return err
	}

	if err := u.enforcer.SavePolicy(); err != nil {
		u.logger.Error("保存策略失败", zap.Error(err))
		return err
	}

	u.logger.Info("删除角色权限成功", zap.String("role", role))
	return nil
}

func (u *CasbinUsecase) AddPolicies(policies [][]string) (bool, error) {
	success, err := u.enforcer.AddPolicies(policies)
	if err != nil {
		u.logger.Error("批量添加权限失败", zap.Any("policies", policies), zap.Error(err))
		return false, err
	}
	if !success {
		u.logger.Warn("批量添加权限未成功，可能权限已存在", zap.Any("policies", policies))
		return false, nil
	}

	if err := u.enforcer.SavePolicy(); err != nil {
		u.logger.Error("保存策略失败", zap.Error(err))
		return false, err
	}

	u.logger.Info("批量添加权限成功", zap.Any("policies", policies))
	return true, nil
}
