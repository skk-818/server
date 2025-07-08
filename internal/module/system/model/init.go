package model

type Init struct {
	BaseModel
	Name        string `gorm:"type:varchar(64);unique;not null;comment:模块名称"`
	Initialized int64  `gorm:"not null;default:0;comment:是否已初始化"`
	Version     string `gorm:"type:varchar(32);default:'';comment:初始化版本"`
	Description string `gorm:"type:varchar(255);default:'';comment:备注信息"`
}

func (m *Init) IsInitialized() bool {
	return m.Initialized == InitInitialized
}

func (m *Init) TableName() string {
	return "sys_init"
}

const (
	InitInitialized    = 1
	InitNotInitialized = 0
	InitNameRole       = "Role"
	InitNameUser       = "User"
	InitNameMenu       = "Menu"
	InitNameCasbin     = "Casbin"
)

var InitCol = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	Name        string
	Initialized string
	Version     string
	Description string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Name:        "name",
	Initialized: "initialized",
	Version:     "version",
	Description: "description",
}
