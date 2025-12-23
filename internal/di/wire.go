// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"server/internal/core"
	"server/internal/core/server"
	"server/internal/middleware"
	imApi "server/internal/module/im/api"
	systemApi "server/internal/module/system/api"
	"server/internal/module/system/biz"
	"server/internal/module/system/repo"
	"server/internal/router"

	wire "github.com/google/wire"
)

func InitApp(path string) (*server.HTTPServer, error) {
	wire.Build(
		core.ProviderSet,
		router.ProviderSet,
		middleware.ProviderSet,
		systemApi.ProviderSet,
		imApi.ProviderSet,
		biz.ProviderSet,
		repo.ProviderSet,
	)

	return nil, nil
}
