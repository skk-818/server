package biz

import (
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
)

type MenuUsecase struct {
	logger   logger.Logger
	menuRepo repo.MenuRepo
}

func NewMenuUsecase(logger logger.Logger, menuRepo repo.MenuRepo) *MenuUsecase {
	return &MenuUsecase{
		logger:   logger,
		menuRepo: menuRepo,
	}
}
