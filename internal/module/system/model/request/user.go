package request

type UserListReq struct {
	Username *string `json:"username"`
	Status   *int    `json:"status"`
	Gender   *int    `json:"gender"`
	PageInfo
}
