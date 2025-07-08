package usecase

import (
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/usecase/repo"
	"sync"
)

type (
	CasbinUsecase struct {
		logger   logger.Logger
		enforcer *casbin.Enforcer
	}

	casbinUsecase interface {
		HasPolicy([]string) (bool, error)
		AddPolicy([]string) (bool, error)
	}
)

var (
	once     sync.Once
	enforcer *casbin.Enforcer
)

func NewCasbinUsecase(logger logger.Logger, repo repo.CasbinRepo) (*CasbinUsecase, error) {
	var err error
	once.Do(func() {
		logger.Info("开始初始化 Casbin Enforcer...")
		db := repo.AdapterDB()
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
	}, nil
}

func (u *CasbinUsecase) Enforce(sub, obj, act string) (bool, error) {
	return u.enforcer.Enforce(sub, obj, act)
}

func (u *CasbinUsecase) HasPolicy(policy []string) (bool, error) {
	return u.enforcer.HasPolicy(policy)
}

func (u *CasbinUsecase) AddPolicy(policy []string) (bool, error) {
	return u.enforcer.AddPolicy(policy)
}
