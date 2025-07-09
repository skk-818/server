// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	wire "github.com/google/wire"
	"server/internal/core"
	"server/internal/core/server"
	"server/internal/middleware"
	"server/internal/module/system/api"
	"server/internal/module/system/repo"
	"server/internal/module/system/usecase"
	"server/internal/router"
)

func InitApp() (*server.HTTPServer, error) {
	wire.Build(
		core.ProviderSet,
		router.ProviderSet,
		middleware.ProviderSet,
		api.ProviderSet,
		usecase.ProviderSet,
		repo.ProviderSet,
	)

	return nil, nil
}
