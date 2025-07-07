package usecase

import (
	"context"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model"
	"server/internal/module/system/usecase/repo"
)

type InitUsecase struct {
	logger        logger.Logger
	initRepo      repo.InitRepo
	userRepo      repo.UserRepo
	roleRepo      repo.RoleRepo
	casbinUsecase casbinUsecase
}

func NewInitUsecase(
	logger logger.Logger,
	initRepo repo.InitRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
	casbinUsecase casbinUsecase,
) *InitUsecase {
	return &InitUsecase{
		logger:        logger,
		initRepo:      initRepo,
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		casbinUsecase: casbinUsecase,
	}
}

func (u *InitUsecase) InitIfNeeded() error {
	if !u.RoleIsInitialized() {
		if err := u.RoleInitialize(); err != nil {
			u.logger.Error("初始化角色失败", zap.Any("err", err))
			return err
		}
	}
	u.logger.Info("role initialized")

	if !u.UserIsInitialized() {
		if err := u.UserInitialize(); err != nil {
			u.logger.Error("初始化用户失败", zap.Any("err", err))
			return err
		}
	}
	u.logger.Info("user initialized")

	if !u.MenuIsInitialized() {
		if err := u.MenuInitialize(); err != nil {
			u.logger.Error("初始化菜单失败", zap.Any("err", err))
			return err
		}
	}
	u.logger.Info("menu initialized")

	if !u.CasbinIsInitialized() {
		if err := u.CasbinInitialize(); err != nil {
			u.logger.Error("初始化用户失败", zap.Any("err", err))
			return err
		}
	}
	u.logger.Info("casbin initialized")

	return nil
}

func (u *InitUsecase) RoleIsInitialized() bool {
	initialized, err := u.initRepo.IsInitialized("role")
	if err != nil {
		u.logger.Error("查询系统初始化状态失败", zap.Any("err", err))
		return false
	}

	return initialized
}

func (u *InitUsecase) UserIsInitialized() bool {
	initialized, err := u.initRepo.IsInitialized("user")
	if err != nil {
		u.logger.Error("查询系统初始化状态失败", zap.Any("err", err))
		return false
	}

	return initialized
}

func (u *InitUsecase) MenuIsInitialized() bool {
	initialized, err := u.initRepo.IsInitialized("user")
	if err != nil {
		u.logger.Error("查询系统初始化状态失败", zap.Any("err", err))
		return false
	}

	return initialized
}

func (u *InitUsecase) CasbinIsInitialized() bool {
	initialized, err := u.initRepo.IsInitialized("user")
	if err != nil {
		u.logger.Error("查询系统初始化状态失败", zap.Any("err", err))
		return false
	}

	return initialized
}

func (u *InitUsecase) RoleInitialize() error {
	// 初始化管理员角色
	role := &model.Role{
		BaseModel: model.BaseModel{
			ID: 1,
		},
		Name:      "超级管理员",
		Key:       "R_ADMIN",
		Status:    1,
		DataScope: "all",
		Sort:      1,
		IsSystem:  1,
		Remark:    "系统初始化超级管理员",
		Users:     nil,
	}

	return u.roleRepo.Create(context.Background(), role)
}

func (u *InitUsecase) UserInitialize() error {
	return nil
}

func (u *InitUsecase) MenuInitialize() error {
	return nil
}

func (u *InitUsecase) CasbinInitialize() error {
	const roleAdmin = "R_ADMIN"

	policy := []string{roleAdmin, "*", "*"}

	exist, err := u.casbinUsecase.HasPolicy(policy)
	if err != nil {
		u.logger.Error("检查 Casbin 策略失败", zap.Any("err", err))
		return err
	}
	if exist {
		u.logger.Info("Casbin 超级管理员策略已存在，跳过初始化")
		return nil
	}

	if ok, err := u.casbinUsecase.AddPolicy(policy); err != nil {
		u.logger.Error("添加 Casbin 超级管理员策略失败", zap.Any("err", err))
		return err
	} else if !ok {
		u.logger.Warn("Casbin 超级管理员策略未添加成功（可能已存在）")
		return nil
	}
	u.logger.Info("Casbin 超级管理员策略添加成功")
	return nil
}
