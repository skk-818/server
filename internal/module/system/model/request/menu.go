package request

type MenuListReq struct {
	PageInfo
	Name     string `json:"name"`
	ParentID uint64 `json:"parentId"`
}
