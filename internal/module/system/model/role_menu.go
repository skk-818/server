package model

type RoleMenu struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID uint64 `gorm:"not null;index;uniqueIndex:idx_role_menu" json:"role_id"`
	MenuID uint64 `gorm:"not null;index;uniqueIndex:idx_role_menu" json:"menu_id"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}
