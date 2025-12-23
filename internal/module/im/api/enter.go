package api

type IMApi struct {
	UserApi    *UserApi
	MessageApi *MessageApi
	GroupApi   *GroupApi
}

func NewIMApi(
	userApi *UserApi,
	messageApi *MessageApi,
	groupApi *GroupApi,
) *IMApi {
	return &IMApi{
		UserApi:    userApi,
		MessageApi: messageApi,
		GroupApi:   groupApi,
	}
}
