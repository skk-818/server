package biz

import (
	"context"
	"fmt"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/pkg"
	"server/pkg/errorx"
	"time"

	"go.uber.org/zap"
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

func (u *AuthUsecase) Login(ctx context.Context, req *request.LoginReq, ip string) (*reply.LoginReply, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, errorx.ErrInternal
	}

	if user == nil {
		return nil, errorx.ErrUserNotFound
	}

	if user.Status != model.UserStatusEnable {
		return nil, errorx.ErrUserDisabled
	}
	if !pkg.CheckPassword(user.Password, req.Password) {
		return nil, errorx.ErrUserPasswordNotMatch
	}

	roleKeys := u.getActiveRoleKeys(user.Roles)
	if len(roleKeys) == 0 {
		return nil, errorx.ErrUserNotRole
	}

	_ = u.userRepo.UpdateLastLogin(ctx, uint(user.ID), ip)

	return u.generateTokens(uint(user.ID), user.Username, roleKeys)
}

func (u *AuthUsecase) Register(ctx context.Context, req *request.RegisterReq) error {
	user, err := u.userRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}
	if user != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername user exist", zap.Any("req", req))
		return errorx.ErrUserConflict
	}

	role, err := u.roleRepo.FindByKey(ctx, model.RoleKeyUser)
	if err != nil {
		u.logger.Error("[AuthUsecase] roleRepo.FindByKey error", zap.Any("req", req), zap.Error(err))
		return errorx.ErrInternal
	}
	if role == nil {
		u.logger.Error("[AuthUsecase] roleRepo.FindByKey role not found", zap.Any("req", req))
		return errorx.ErrRoleNotFound
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
		return errorx.ErrInternal
	}

	return nil
}

func (u *AuthUsecase) EmailLogin(ctx context.Context, req *request.EmailLoginReq, ip string) (*reply.LoginReply, error) {
	// TODO: 验证邮箱验证码

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errorx.ErrInternal
	}

	if user == nil {
		role, err := u.roleRepo.FindByKey(ctx, model.RoleKeyUser)
		if err != nil || role == nil {
			return nil, errorx.ErrRoleNotFound
		}

		user = &model.User{
			Username: fmt.Sprintf("u_%d", time.Now().UnixNano()),
			Email:    req.Email,
			Nickname: fmt.Sprintf("用户%d", time.Now().UnixNano()%1e6),
			Status:   model.UserStatusEnable,
			IsAdmin:  model.UserNotSystem,
			Roles:    []*model.Role{role},
		}

		if err := u.userRepo.Create(ctx, user); err != nil {
			return nil, errorx.ErrInternal
		}
	}

	if user.Status != model.UserStatusEnable {
		return nil, errorx.ErrUserDisabled
	}

	roleKeys := u.getActiveRoleKeys(user.Roles)
	if len(roleKeys) == 0 {
		return nil, errorx.ErrUserNotRole
	}

	_ = u.userRepo.UpdateLastLogin(ctx, uint(user.ID), ip)

	return u.generateTokens(uint(user.ID), user.Username, roleKeys)
}

func (u *AuthUsecase) getActiveRoleKeys(roles []*model.Role) []string {
	roleKeys := make([]string, 0, len(roles))
	for _, role := range roles {
		if role.Status == model.RoleStatusEnable {
			roleKeys = append(roleKeys, role.Key)
		}
	}
	return roleKeys
}

func (u *AuthUsecase) generateTokens(userID uint, username string, roleKeys []string) (*reply.LoginReply, error) {
	accessToken, err := u.jwtUsecase.GenerateAccessToken(userID, username, roleKeys)
	if err != nil {
		return nil, errorx.ErrAuthGenerateTokenFail
	}

	refreshToken, err := u.jwtUsecase.GenerateRefreshToken(userID, username, roleKeys)
	if err != nil {
		return nil, errorx.ErrAuthGenerateTokenFail
	}

	return &reply.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
