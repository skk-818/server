package repo

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/internal/module/system/model"
	"server/internal/module/system/usecase/repo"
)

type initRepo struct {
	db *gorm.DB
}

func NewInitRepo(db *gorm.DB) repo.InitRepo {
	return &initRepo{db: db}
}

func (r *initRepo) IsInitialized(name string) (bool, error) {
	var init model.Init
	err := r.db.Where("name = ?", name).First(&init).Error
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
