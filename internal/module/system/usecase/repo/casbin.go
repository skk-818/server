package repo

import "gorm.io/gorm"

type CasbinRepo interface {
	AdapterDB() *gorm.DB
}
