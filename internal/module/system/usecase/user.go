package usecase

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"server/internal/core/logger"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
	"server/pkg"
	"server/pkg/errorx"
	"strings"
	"time"
)

type UserUsecase struct {
	logger     logger.Logger
	userRepo   repo.UserRepo
	roleRepo   repo.RoleRepo
	jwtUsecase jwtUsecase
}

func NewUserUsecase(
	logger logger.Logger,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
	jwtUsecase jwtUsecase,
) *UserUsecase {
	return &UserUsecase{
		logger:     logger,
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		jwtUsecase: jwtUsecase,
	}
}

func (u *UserUsecase) Detail(ctx context.Context, userId int) (*reply.UserDetailReply, error) {
	user, err := u.userRepo.Find(ctx, int64(userId))
	if err != nil {
		u.logger.Error("[UserUsecase] userRepo.Find err", zap.Any("userId", userId), zap.Error(err))
		return nil, err
	}
	if user == nil {
		u.logger.Warn("[UserUsecase] userRepo.Find user not find", zap.Any("userId", userId), zap.Error(err))
		return nil, errorx.ErrUserNotFound
	}
	return reply.BuilderUserDetailReply(user), nil
}

func (u *UserUsecase) Create(ctx context.Context, req *request.CreateUserReq) error {
	roles, err := u.roleRepo.FindByKeys(ctx, req.RoleKey)
	if err != nil {
		u.logger.Error("[UserUsecase] roleRepo.FindByKeys err", zap.Any("req", req), zap.Error(err))
		return err
	}
	if len(roles) == 0 || len(roles) < len(req.RoleKey) {
		u.logger.Error("[UserUsecase] roleRepo.FindByKeys roles is empty", zap.Any("req", req), zap.Error(err))
		return errorx.ErrRoleNotFound
	}

	createUser := u.randomUser(roles, req.Username, pkg.HashPassword(req.Password))
	err = u.userRepo.Create(ctx, createUser)
	if err != nil {
		u.logger.Error("[UserUsecase] userRepo.Create err", zap.Any("req", req), zap.Error(err))
		return err
	}

	return nil
}

func (u *UserUsecase) randomUser(roles []*model.Role, username, password string) *model.User {
	rand.NewSource(time.Now().UnixNano())

	jobTitles := []string{"开发工程师", "测试工程师", "产品经理", "UI设计师", "运维工程师"}
	positions := []string{"后端开发", "前端开发", "移动开发", "测试", "系统运维"}
	departments := []string{"研发部", "测试部", "产品部", "运维部"}
	tagsList := []string{"积极", "爱摸鱼", "努力", "卷王", "佛系", "摆烂", "天才", "社牛", "内向"}

	randomPick := func(list []string) string {
		return list[rand.Intn(len(list))]
	}

	randomPhone := func() string {
		return fmt.Sprintf("1%09d", rand.Intn(1e9))
	}

	randomEmail := func(username string) string {
		return fmt.Sprintf("%s%d@qq.com", username, rand.Intn(10000))
	}

	nickname := fmt.Sprintf("用户_%d", rand.Intn(10000))

	return &model.User{
		Username:   username,
		Password:   password,
		Nickname:   nickname,
		Email:      randomEmail(username),
		Phone:      randomPhone(),
		Avatar:     "",
		Gender:     model.UserGenderMale, // 也可以 rand.Intn(2)
		Status:     model.UserStatusEnable,
		IsAdmin:    model.UserNotSystem,
		Province:   "四川省",
		City:       "成都市",
		District:   "武侯区",
		Address:    "四川省成都市武侯区某某街道",
		Position:   randomPick(positions),
		Department: randomPick(departments),
		JobTitle:   randomPick(jobTitles),
		Tags:       strings.Join(randomTags(tagsList, 2), ","),
		Roles:      roles,
	}
}

func randomTags(tags []string, n int) []string {
	rand.Shuffle(len(tags), func(i, j int) {
		tags[i], tags[j] = tags[j], tags[i]
	})
	if n > len(tags) {
		n = len(tags)
	}
	return tags[:n]
}

func (u *UserUsecase) Delete(ctx context.Context, req *request.DeleteUserReq) error {
	users, err := u.userRepo.FindByIds(ctx, req.Ids)
	if err != nil {
		u.logger.Error("[UserUsecase] userRepo.FindByIds err", zap.Any("req", req), zap.Error(err))
		return err
	}
	if len(users) == 0 || len(users) < len(req.Ids) {
		u.logger.Error("[UserUsecase] userRepo.FindByIds users is empty", zap.Any("req", req), zap.Error(err))
		return errorx.ErrUserNotFound
	}

	var deleteUserIds []int64
	for _, user := range users {
		deleteUserIds = append(deleteUserIds, int64(user.ID))
	}

	if len(deleteUserIds) < len(req.Ids) {
		u.logger.Error("[UserUsecase] userRepo.FindByIds users is empty", zap.Any("req", req), zap.Error(err))
		return errorx.ErrUserIsSystem
	}

	err = u.userRepo.BatchDelete(ctx, deleteUserIds)
	if err != nil {
		u.logger.Error("[UserUsecase] userRepo.Delete err", zap.Any("req", req), zap.Error(err))
		return err
	}

	return nil
}

func (u *UserUsecase) Login(ctx context.Context, req *request.LoginReq) (*reply.LoginReply, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername error", zap.Any("req", req), zap.Error(err))
		return nil, errorx.ErrInternal
	}
	if user == nil {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername user not find", zap.Any("req", req))
		return nil, errorx.ErrUserDisabled
	}
	if user.Status != model.UserStatusEnable {
		u.logger.Warn("[AuthUsecase] user not enable", zap.Any("req", req))
		return nil, err
	}

	if !pkg.CheckPassword(user.Password, req.Password) {
		u.logger.Error("[AuthUsecase] userRepo.FindByUsername password not match", zap.Any("req", req))
		return nil, errorx.ErrUserPasswordNotMatch
	}

	roleKeys := make([]string, 0)
	if len(user.Roles) > 0 {
		for i := range user.Roles {
			if user.Roles[i].Status == model.RoleStatusEnable { // 添加状态开启的 role
				roleKeys = append(roleKeys, user.Roles[i].Name)
			}
		}
	}
	if len(roleKeys) == 0 {
		return nil, errorx.ErrUserNotRole
	}

	accessToken, err := u.jwtUsecase.GenerateAccessToken(uint(user.ID), user.Username, roleKeys)
	if err != nil {
		u.logger.Error("[AuthUsecase] GenerateAccessToken error", zap.Any("req", req), zap.Error(err))
		return nil, errorx.ErrAuthGenerateTokenFail
	}

	refreshToken, err := u.jwtUsecase.GenerateRefreshToken(uint(user.ID), user.Username, roleKeys)
	if err != nil {
		u.logger.Error("[AuthUsecase] GenerateRefreshToken error", zap.Any("req", req), zap.Error(err))
		return nil, errorx.ErrAuthGenerateTokenFail
	}

	return &reply.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserUsecase) Register(ctx context.Context, req *request.RegisterReq) error {
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
