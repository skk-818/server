package usecase

var ProviderSet = wire.NewSet(
	NewUserUsecase,
)
