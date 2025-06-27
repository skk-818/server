package model

type UserRole struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint64 `gorm:"not null;index;uniqueIndex:idx_user_role" json:"user_id"`
	RoleID uint64 `gorm:"not null;index;uniqueIndex:idx_user_role" json:"role_id"`
}

func (m *UserRole) TableName() string {
	return "sys_user_role"
}
