package model

type Role struct {
	BaseModel
	Name      string `gorm:"size:64;not null;comment:角色名称" json:"name"`
	Code      string `gorm:"size:64;uniqueIndex;not null;comment:角色编码（唯一英文标识）" json:"code"`
	Status    int64  `gorm:"type:tinyint(1);default:1;not null;comment:角色状态（1启用，0禁用）" json:"status"`
	DataScope string `gorm:"size:32;default:'all';not null;comment:数据权限范围（all=全部，dept=本部门，self=本人）" json:"dataScope"`
	Sort      int64  `gorm:"default:0;not null;comment:显示顺序（越小越靠前）" json:"sort"`
	IsSystem  int64  `gorm:"type:tinyint(1);default:0;not null;comment:是否为系统内置角色（1是 0否）" json:"isSystem"`
	Remark    string `gorm:"size:255;default:'';not null;comment:备注信息" json:"remark"`

	Users []*User `gorm:"many2many:sys_user_role;" json:"users"`
}

func (m *Role) TableName() string {
	return "sys_role"
}

const (
	RoleStatusEnable  = 1
	RoleStatusDisable = 0
)

var RoleCol = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	Name      string
	Code      string
	Status    string
	DataScope string
	Sort      string
	IsSystem  string
	Remark    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Name:      "name",
	Code:      "code",
	Status:    "status",
	DataScope: "data_scope",
	Sort:      "sort",
	IsSystem:  "is_system",
	Remark:    "remark",
}
