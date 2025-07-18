package api

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewSystemApi,
	NewUserApi,
	NewAuthApi,
	NewRoleApi,
	NewApiApi,
	NewMenuApi,
)
