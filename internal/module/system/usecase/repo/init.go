package repo

import "gorm.io/gorm/schema"

type InitRepo interface {
	AutoMigrate([]schema.Tabler) error
	IsInitialized(string) (bool, error)
	SetInitialized(string, string, string) error
}
