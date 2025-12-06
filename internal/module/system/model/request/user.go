package request

type UserListReq struct {
	Username *string `json:"username" form:"username"`
	Status   *int    `json:"status" form:"status"`
	Gender   *int    `json:"gender" form:"gender"`
	PageInfo
}

type CreateUserReq struct {
	Username string   `json:"username" validate:"required"`
	Password string   `json:"password" validate:"required"`
	RoleKey  []string `json:"roleKey" validate:"required,gt=0"`
}

type DeleteUserReq struct {
	Ids []int64 `json:"ids" validate:"required,gt=0"`
}
