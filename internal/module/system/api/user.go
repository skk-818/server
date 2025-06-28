package api

import (
	"server/internal/core/config"
	"server/internal/module/system/service"
)

type UserApi struct {
	config      *config.Config
	userService *service.UserService
}

func NewUserApi(config *config.Config, userService *service.UserService) *UserApi {
	return &UserApi{
		config: config,
	}
}
