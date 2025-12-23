package request

type RoleListReq struct {
	PageInfo
	Name string `json:"name" form:"name"`
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

// 分配API权限请求
type AssignApiPermissionsReq struct {
	RoleId int64   `json:"roleId" validate:"required,notzero"` // 角色ID
	ApiIds []int64 `json:"apiIds" validate:"required"`         // API ID列表
}

// 分配菜单权限请求
type AssignMenuPermissionsReq struct {
	RoleId          int64    `json:"roleId" validate:"required,notzero"` // 角色ID
	MenuIds         []uint64 `json:"menuIds" validate:"required"`        // 菜单ID列表
	HalfCheckedKeys []uint64 `json:"halfCheckedKeys"`                    // 半选中的节点（父节点）
}
