package usecase

import (
	"server/internal/core/logger"
	"server/internal/module/system/usecase/repo"
)

type RoleUsecase struct {
	logger   logger.Logger
	roleRepo repo.RoleRepo
}

func NewRoleUsecase(
	logger logger.Logger,
	roleRepo repo.RoleRepo,
) *RoleUsecase {
	return &RoleUsecase{
		logger:   logger,
		roleRepo: roleRepo,
	}
}
