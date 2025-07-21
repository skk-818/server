package request

type RoleListReq struct {
	PageInfo
	Name string `json:"name"`
}

type CreateRoleReq struct {
	Name   string `json:"name" validate:"required"`
	Key    string `json:"key" validate:"required"`
	Remark string `json:"remark" validate:"required"`
	Sort   *int64 `json:"sort" validate:"required,notzero"`
}

type DeleteRoleReq struct {
	Ids []int64 `json:"ids" validate:"required,min=1,dive,gt=0"`
}

type UpdateRoleReq struct {
	Id     *int64 `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Key    string `json:"key" validate:"required"`
	Remark string `json:"remark" validate:"required"`
	Sort   *int64 `json:"sort" validate:"required,notzero"`
}

type GetRoleReq struct {
	Id *int64 `json:"id" validate:"required,notzero"`
}
