package service

var ProviderSet = wire.NewSet(
	NewUserService,
)
