package request

type RoleListReq struct {
	PageInfo
	Name   string `json:"name"`
	Status *int64 `json:"status"`
}
