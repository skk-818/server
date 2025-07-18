package usecase

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
	"server/pkg"
	"server/pkg/xerror"
	"time"
)

type AuthUsecase struct {
	logger     logger.Logger
	userRepo   repo.UserRepo
	roleRepo   repo.RoleRepo
	jwtUsecase jwtUsecase
}

func NewAuthUsecase(logger logger.Logger, userRepo repo.UserRepo, roleRepo repo.RoleRepo, jwtUsecase jwtUsecase) *AuthUsecase {
	return &AuthUsecase{
		logger:     logger,
		roleRepo:   roleRepo,
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

func (u *AuthUsecase) Register(ctx context.Context, req *request.RegisterReq) error {
	user, err := u.userRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername error", zap.Any("req", req), zap.Error(err))
		return xerror.ErrInternal
	}
	if user != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername user exist", zap.Any("req", req))
		return xerror.ErrUserConflict
	}

	role, err := u.roleRepo.FindByKey(ctx, model.RoleKeyUser)
	if err != nil {
		u.logger.Error("[AuthUsecase] roleRepo.FindByKey error", zap.Any("req", req), zap.Error(err))
		return xerror.ErrInternal
	}
	if role == nil {
		u.logger.Error("[AuthUsecase] roleRepo.FindByKey role not found", zap.Any("req", req))
		return xerror.ErrRoleNotFound
	}

	createUser := &model.User{
		Username: fmt.Sprintf("u_%d", time.Now().UnixNano()),
		Password: pkg.HashPassword(req.Password),
		Nickname: fmt.Sprintf("用户%d", time.Now().UnixNano()%1e6),
		Phone:    req.Phone,
		Avatar:   "https://cdn.example.com/avatar/default.png",
		Status:   model.UserStatusEnable,
		IsAdmin:  model.UserNotSystem,
		Position: "普通用户",
		Tags:     "新注册",
		Roles:    []*model.Role{role},
	}

	if err := u.userRepo.Create(ctx, createUser); err != nil {
		u.logger.Error("[AuthUsecase] userRepo.Create error", zap.Any("user", createUser), zap.Error(err))
		return xerror.ErrInternal
	}

	return nil
}
