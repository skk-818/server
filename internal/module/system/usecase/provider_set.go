package usecase

import (
	"github.com/google/wire"
	"server/internal/middleware"
)

var ProviderSet = wire.NewSet(
	NewUserUsecase,
	NewJwtUsecase,
	wire.Bind(new(middleware.JwtParse), new(*JwtUsecase)),
)
