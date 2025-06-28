package service

import "server/internal/module/system/usecase"

type UserService struct {
	userUsecase *usecase.UserUsecase
}

func NewUserService(userUsecase *usecase.UserUsecase) *UserService {
	return &UserService{
		userUsecase: userUsecase,
	}
}
