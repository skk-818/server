package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/response"
	"strings"
)

type MenuUsecase struct {
	logger       logger.Logger
	menuRepo     repo.MenuRepo
	roleMenuRepo repo.RoleMenuRepo
	userRepo     repo.UserRepo
}

func NewMenuUsecase(logger logger.Logger, menuRepo repo.MenuRepo, roleMenuRepo repo.RoleMenuRepo, userRepo repo.UserRepo) *MenuUsecase {
	return &MenuUsecase{
		logger:       logger,
		menuRepo:     menuRepo,
		roleMenuRepo: roleMenuRepo,
		userRepo:     userRepo,
	}
}

func (u *MenuUsecase) GetAllMenuTree(ctx context.Context) ([]*response.MenuTreeResp, error) {
	menus, err := u.menuRepo.GetAllEnabled(ctx)
	if err != nil {
		return nil, err
	}
	return u.buildMenuTree(menus, 0), nil
}

func (u *MenuUsecase) GetMenuTree(ctx context.Context) ([]*response.MenuTreeResp, error) {
	// 获取所有启用的菜单
	menus, err := u.menuRepo.GetAllEnabled(ctx)
	if err != nil {
		return nil, err
	}

	// 从 context 中获取用户ID
	userIDVal := ctx.Value("userID")
	if userIDVal == nil {
		// 如果没有用户ID，返回空菜单（未登录或token无效）
		return []*response.MenuTreeResp{}, nil
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		return []*response.MenuTreeResp{}, nil
	}

	// 获取用户信息（包含角色）
	user, err := u.userRepo.Find(ctx, int64(userID))
	if err != nil || user == nil {
		return []*response.MenuTreeResp{}, nil
	}

	// 获取用户所有角色的菜单权限
	allowedMenuIDs := make(map[uint64]bool)
	for _, role := range user.Roles {
		menuIDs, err := u.roleMenuRepo.GetMenuIdsByRoleId(ctx, uint64(role.ID))
		if err != nil {
			continue
		}
		for _, menuID := range menuIDs {
			allowedMenuIDs[menuID] = true
		}
	}

	// 收集所有需要显示的菜单ID（包括父菜单）
	menuIDsToShow := make(map[uint64]bool)
	for menuID := range allowedMenuIDs {
		menuIDsToShow[menuID] = true
		// 递归添加所有父菜单
		u.addParentMenus(menus, menuID, menuIDsToShow)
	}

	// 过滤菜单：保留用户有权限的菜单及其父菜单
	var filteredMenus []*model.Menu
	for _, menu := range menus {
		if menuIDsToShow[menu.ID] {
			filteredMenus = append(filteredMenus, menu)
		}
	}

	return u.buildMenuTree(filteredMenus, 0), nil
}

// addParentMenus 递归添加父菜单ID
func (u *MenuUsecase) addParentMenus(menus []*model.Menu, menuID uint64, result map[uint64]bool) {
	for _, menu := range menus {
		if menu.ID == menuID && menu.ParentID != 0 {
			result[menu.ParentID] = true
			u.addParentMenus(menus, menu.ParentID, result)
			break
		}
	}
}

func (u *MenuUsecase) Create(ctx context.Context, req *model.Menu) error {
	return u.menuRepo.Create(ctx, req)
}

func (u *MenuUsecase) Update(ctx context.Context, req *model.Menu) error {
	return u.menuRepo.Update(ctx, req)
}

func (u *MenuUsecase) Delete(ctx context.Context, id int64) error {
	return u.menuRepo.Delete(ctx, id)
}

func (u *MenuUsecase) List(ctx context.Context, req *model.Menu) (*reply.ListMenuReply, error) {
	menus, err := u.menuRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return reply.BuilderListMenuReply(menus), nil
}

func (u *MenuUsecase) buildMenuTree(menus []*model.Menu, parentID uint64) []*response.MenuTreeResp {
	var tree []*response.MenuTreeResp
	for _, menu := range menus {
		if menu.ParentID == parentID {
			node := &response.MenuTreeResp{
				ID:        menu.ID,
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Redirect:  menu.Redirect,
				Meta: response.MenuMeta{
					Title:         menu.Title,
					Icon:          menu.Icon,
					IsHide:        menu.Hidden == 1,
					IsHideTab:     menu.HideTab == 1,
					Link:          menu.Link,
					IsIframe:      menu.IsIframe == 1,
					KeepAlive:     menu.KeepAlive == 1,
					FixedTab:      menu.FixedTab == 1,
					ShowBadge:     menu.ShowBadge == 1,
					ShowTextBadge: menu.TextBadge,
					ActivePath:    menu.ActivePath,
					IsFullPage:    menu.FullPage == 1,
				},
			}
			if menu.Roles != "" {
				node.Meta.Roles = strings.Split(menu.Roles, ",")
			}
			node.Children = u.buildMenuTree(menus, menu.ID)
			tree = append(tree, node)
		}
	}
	return tree
}
