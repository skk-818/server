package usecase

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model"
	"server/internal/module/system/usecase/repo"
	"server/pkg/xerror"
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

func (u *UserUsecase) GetUserInfo(c *gin.Context, userId int) (*model.User, error) {
	user, err := u.userRepo.Find(c, int64(userId))
	if err != nil {
		u.logger.Error("[UserUsecase] userRepo.Find err", zap.Any("userId", userId), zap.Error(err))
		return nil, err
	}
	if user == nil {
		u.logger.Warn("[UserUsecase] userRepo.Find user not find", zap.Any("userId", userId), zap.Error(err))
		return nil, errorx.ErrUserNotFound
	}
	return user, nil
}
