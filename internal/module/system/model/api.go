package model

type Api struct {
	BaseModel
	Name        string `gorm:"size:128;not null;comment:接口名称，比如 用户列表接口" json:"name"`
	Path        string `gorm:"size:256;not null;uniqueIndex:uk_path_method;comment:接口路径，比如 /api/user" json:"path"`
	Method      string `gorm:"size:16;not null;uniqueIndex:uk_path_method;comment:请求方法，比如 GET、POST" json:"method"`
	Description string `gorm:"size:512;not null;default:'';comment:接口描述" json:"description"`
	Group       string `gorm:"size:64;not null;default:'';comment:接口分组，比如用户管理、订单管理" json:"group"`
	Status      int64  `gorm:"type:tinyint(1);not null;default:1;comment:状态（1启用，0禁用）" json:"status"`
}

func (m *Api) TableName() string {
	return "sys_api"
}

var ApiCol = struct {
	ID          string
	Name        string
	Path        string
	Method      string
	Description string
	Group       string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Path:        "path",
	Method:      "method",
	Description: "description",
	Group:       "group",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}
