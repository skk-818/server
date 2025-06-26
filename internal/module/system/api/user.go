package api

import "server/internal/core/config"

type UserApi struct {
	config *config.Config
}

func NewUserApi(config *config.Config) *UserApi {
	return &UserApi{
		config: config,
	}
}
