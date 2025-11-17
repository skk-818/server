package biz

import (
	"server/internal/middleware"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewJwtUsecase,
	wire.Bind(new(middleware.JwtParse), new(*JwtUsecase)),
	wire.Bind(new(jwtUsecase), new(*JwtUsecase)),
	NewCasbinUsecase,
	wire.Bind(new(middleware.CabinEnforce), new(*CasbinUsecase)),
	wire.Bind(new(casbinUsecase), new(*CasbinUsecase)),
	NewInitUsecase,
	NewCronUsecase,

	NewUserUsecase,
	NewAuthUsecase,
	NewRoleUsecase,
	NewApiUsecase,
	NewMenuUsecase,
)
