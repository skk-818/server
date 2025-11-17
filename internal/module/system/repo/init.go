package repo

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
)

type initRepo struct {
	db *gorm.DB
}

func NewInitRepo(db *gorm.DB) repo.InitRepo {
	return &initRepo{db: db}
}

func (r *initRepo) AutoMigrate(tables []schema.Tabler) error {
	for _, table := range tables {
		if err := r.db.AutoMigrate(table); err != nil {
			return err
		}
	}
	return nil
}

func (r *initRepo) IsInitialized(name string) (bool, error) {
	var init model.Init
	err := r.db.Where(model.InitCol.Name+" = ?", name).First(&init).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil //
		}
		return false, err
	}
	return init.IsInitialized(), nil
}

func (r *initRepo) SetInitialized(name, version, description string) error {
	initRecord := model.Init{
		Name:        name,
		Initialized: model.InitInitialized,
		Version:     version,
		Description: description,
	}

	return r.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"initialized", "version", "description", "updated_at"}),
		},
	).Create(&initRecord).Error
}
