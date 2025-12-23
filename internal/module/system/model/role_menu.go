package model

// RoleMenu 角色菜单关联表
type RoleMenu struct {
	RoleID uint64 `gorm:"primaryKey;not null;comment:角色ID" json:"roleId"`
	MenuID uint64 `gorm:"primaryKey;not null;comment:菜单ID" json:"menuId"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}
