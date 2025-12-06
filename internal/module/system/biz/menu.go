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
	logger   logger.Logger
	menuRepo repo.MenuRepo
}

func NewMenuUsecase(logger logger.Logger, menuRepo repo.MenuRepo) *MenuUsecase {
	return &MenuUsecase{
		logger:   logger,
		menuRepo: menuRepo,
	}
}

func (u *MenuUsecase) GetMenuTree(ctx context.Context) ([]*response.MenuTreeResp, error) {
	menus, err := u.menuRepo.GetAllEnabled(ctx)
	if err != nil {
		return nil, err
	}
	return u.buildMenuTree(menus, 0), nil
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
