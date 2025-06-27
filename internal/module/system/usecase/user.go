package usecase

import (
	"server/internal/core/logger"
	"server/internal/repo"
)

type UserUsecase struct {
	logger   logger.Logger
	userRepo repo.UserRepo
}

func NewUserUsecase(logger logger.Logger, userRepo repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		logger:   logger,
		userRepo: userRepo,
	}
}
