package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/internal/core/mysql"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
)

type roleMenuRepo struct {
	db *gorm.DB
}

func NewRoleMenuRepo(systemDB *mysql.SystemDB) repo.RoleMenuRepo {
	return &roleMenuRepo{db: systemDB.DB}
}

func (r *roleMenuRepo) AssignMenus(ctx context.Context, roleId uint64, menuIds []uint64) error {
	// 先删除该角色的所有菜单权限
	if err := r.DeleteByRoleId(ctx, roleId); err != nil {
		return err
	}

	// 如果没有菜单ID，直接返回
	if len(menuIds) == 0 {
		return nil
	}

	// 批量插入新的菜单权限
	var roleMenus []model.RoleMenu
	for _, menuId := range menuIds {
		roleMenus = append(roleMenus, model.RoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		})
	}

	err := r.db.WithContext(ctx).Create(&roleMenus).Error
	return errors.WithStack(err)
}

func (r *roleMenuRepo) GetMenuIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	var roleMenus []model.RoleMenu
	err := r.db.WithContext(ctx).
		Where("role_id = ?", roleId).
		Find(&roleMenus).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var menuIds []uint64
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuID)
	}

	return menuIds, nil
}

func (r *roleMenuRepo) DeleteByRoleId(ctx context.Context, roleId uint64) error {
	err := r.db.WithContext(ctx).
		Where("role_id = ?", roleId).
		Delete(&model.RoleMenu{}).Error
	return errors.WithStack(err)
}
