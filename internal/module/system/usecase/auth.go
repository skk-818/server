package usecase

import (
	"context"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
	"server/pkg"
	"server/pkg/xerror"
)

type AuthUsecase struct {
	logger     logger.Logger
	userRepo   repo.UserRepo
	jwtUsecase jwtUsecase
}

func NewAuthUsecase(logger logger.Logger, userRepo repo.UserRepo, jwtUsecase jwtUsecase) *AuthUsecase {
	return &AuthUsecase{
		logger:     logger,
		userRepo:   userRepo,
		jwtUsecase: jwtUsecase,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, req *request.LoginReq) (*reply.LoginReply, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername error", zap.Any("req", req), zap.Error(err))
		return nil, xerror.ErrInternal
	}
	if user == nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername user not find", zap.Any("req", req))
		return nil, xerror.ErrUserNotFound
	}

	if !pkg.CheckPassword(user.Password, req.Password) {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername password not match", zap.Any("req", req))
		return nil, xerror.ErrUserPasswordNotMatch
	}

	roleKeys := make([]string, 0)
	if len(user.Roles) > 0 {
		for i := range user.Roles {
			roleKeys = append(roleKeys, user.Roles[i].Key)
		}
	}

	accessToken, err := u.jwtUsecase.GenerateAccessToken(uint(user.ID), user.Username, roleKeys)
	if err != nil {
		u.logger.Error("[AuthUsecase] GenerateAccessToken error", zap.Any("req", req), zap.Error(err))
		return nil, xerror.ErrAuthGenerateTokenFail
	}

	refreshToken, err := u.jwtUsecase.GenerateRefreshToken(uint(user.ID), user.Username, roleKeys)
	if err != nil {
		u.logger.Error("[AuthUsecase] GenerateRefreshToken error", zap.Any("req", req), zap.Error(err))
		return nil, xerror.ErrAuthGenerateTokenFail
	}

	return &reply.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
