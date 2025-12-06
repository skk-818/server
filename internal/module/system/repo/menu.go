package repo

import (
	"context"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/internal/module/system/model/request"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) repo.MenuRepo {
	return &menuRepo{db: db}
}

func (m *menuRepo) Create(ctx context.Context, menu *model.Menu) error {
	err := m.db.WithContext(ctx).Create(menu).Error
	return errors.WithStack(err)
}

func (m *menuRepo) Update(ctx context.Context, menu *model.Menu) error {
	err := m.db.WithContext(ctx).Save(menu).Error
	return errors.WithStack(err)
}

func (m *menuRepo) Delete(ctx context.Context, id int64) error {
	err := m.db.WithContext(ctx).Delete(&model.Menu{}, id).Error
	return errors.WithStack(err)
}

func (m *menuRepo) Find(ctx context.Context, id int64) (*model.Menu, error) {
	var menu model.Menu
	err := m.db.WithContext(ctx).First(&menu, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没找到返回nil,nil
		}
		return nil, errors.WithStack(err)
	}
	return &menu, nil
}

func (m *menuRepo) List(ctx context.Context, req *request.MenuListReq) ([]*model.Menu, int64, error) {
	var menus []*model.Menu
	var total int64

	db := m.db.WithContext(ctx).Model(&model.Menu{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.ParentID != 0 {
		db = db.Where("parent_id = ?", req.ParentID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	offset, limit := req.BuilderOffsetAndLimit()
	if err := db.Order("sort ASC").Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return menus, total, nil
}

func (m *menuRepo) BatchDelete(ctx context.Context, ids []int64) error {
	err := m.db.WithContext(ctx).Where("id IN ?", ids).Delete(&model.Menu{}).Error
	return errors.WithStack(err)
}

func (m *menuRepo) GetAllEnabled(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := m.db.WithContext(ctx).Where("status = ?", 1).Order("sort ASC").Find(&menus).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return menus, nil
}

func (m *menuRepo) GetAll(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := m.db.WithContext(ctx).Order("id ASC").Find(&menus).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return menus, nil
}
