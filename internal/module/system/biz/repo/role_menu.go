package repo

import (
	"context"
)

type RoleMenuRepo interface {
	// AssignMenus 为角色分配菜单权限
	AssignMenus(ctx context.Context, roleId uint64, menuIds []uint64) error
	// GetMenuIdsByRoleId 获取角色的菜单ID列表
	GetMenuIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error)
	// DeleteByRoleId 删除角色的所有菜单权限
	DeleteByRoleId(ctx context.Context, roleId uint64) error
}
