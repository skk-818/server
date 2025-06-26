// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	wire "github.com/google/wire"
	"server/internal/core"
	"server/internal/core/server"
	"server/internal/module/system/api"
	"server/internal/router"
)

func InitApp() (*server.HTTPServer, error) {
	wire.Build(
		core.ProviderSet,
		router.ProviderSet,
		api.ProviderSet,
	)

	return nil, nil
}
