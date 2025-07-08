package model

import "time"

type User struct {
	BaseModel

	Username string `gorm:"size:64;not null;uniqueIndex;comment:用户名" json:"username"`
	Password string `gorm:"size:128;not null;comment:密码" json:"-"`
	Nickname string `gorm:"size:64;not null;default:'';comment:用户昵称" json:"nickname"`
	Email    string `gorm:"size:128;not null;default:'';comment:邮箱" json:"email"`
	Phone    string `gorm:"size:20;not null;default:'';comment:手机号" json:"phone"`
	Avatar   string `gorm:"size:255;not null;default:'';comment:头像URL" json:"avatar"`

	Gender  int64 `gorm:"type:tinyint(1);not null;default:0;comment:性别（0未知 1男 2女）" json:"gender"`
	Status  int64 `gorm:"type:tinyint(1);not null;default:1;comment:账号状态（0禁用 1启用）" json:"status"`
	IsAdmin int64 `gorm:"type:tinyint(1);not null;default:0;comment:是否管理员" json:"isAdmin"`

	// 📍 地理信息
	Province string `gorm:"size:64;not null;default:'';comment:省份" json:"province"`
	City     string `gorm:"size:64;not null;default:'';comment:城市" json:"city"`
	District string `gorm:"size:64;not null;default:'';comment:区县" json:"district"`
	Address  string `gorm:"size:255;not null;default:'';comment:详细地址" json:"address"`

	// 💼 职业信息
	Position   string `gorm:"size:64;not null;default:'';comment:岗位（职务）" json:"position"`
	Department string `gorm:"size:64;not null;default:'';comment:部门" json:"department"`
	JobTitle   string `gorm:"size:64;not null;default:'';comment:职业头衔/职位" json:"jobTitle"`

	// 🏷️ 标签信息
	Tags string `gorm:"size:255;not null;default:'';comment:用户标签（英文逗号分隔）" json:"tags"`

	LastLoginAt *time.Time `gorm:"comment:最后登录时间" json:"lastLoginAt"`
	LastLoginIP string     `gorm:"size:45;not null;default:'';comment:最后登录IP" json:"lastLoginIP"`

	Roles []*Role `gorm:"many2many:sys_user_role;" json:"roles"`
}

func (m *User) TableName() string {
	return "sys_user"
}

const (
	UserStatusEnable  = 1
	UserStatusDisable = 0
	UserGenderMale    = 1
	UserGenderFemale  = 2
	UserIsSystem      = 1
	UserNotSystem     = 0
)

var UserCol = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Username    string
	Password    string
	Nickname    string
	Email       string
	Phone       string
	Avatar      string
	Gender      string
	Status      string
	IsAdmin     string
	Province    string
	City        string
	District    string
	Address     string
	Position    string
	Department  string
	JobTitle    string
	Tags        string
	LastLoginAt string
	LastLoginIP string
	Roles       string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Username:    "username",
	Password:    "password",
	Nickname:    "nickname",
	Email:       "email",
	Phone:       "phone",
	Avatar:      "avatar",
	Gender:      "gender",
	Status:      "status",
	IsAdmin:     "is_admin",
	Province:    "province",
	City:        "city",
	District:    "district",
	Address:     "address",
	Position:    "position",
	Department:  "department",
	JobTitle:    "job_title",
	Tags:        "tags",
	LastLoginAt: "last_login_at",
	LastLoginIP: "last_login_ip",
	Roles:       "Roles",
}
