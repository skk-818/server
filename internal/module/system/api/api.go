package api

type SystemApi struct {
	userApi *UserApi
}

func NewSystemApi() *SystemApi {
	return &SystemApi{}
}
