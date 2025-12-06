package reply

import (
	"server/internal/module/system/model"
	"strings"
)

type MenuMeta struct {
	Title string   `json:"title"`
	Icon  string   `json:"icon,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type MenuReply struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Path      string       `json:"path"`
	Component string       `json:"component,omitempty"`
	Meta      *MenuMeta    `json:"meta"`
	Status    int          `json:"status"`
	UpdatedAt string       `json:"updatedAt"`
	Children  []*MenuReply `json:"children,omitempty"`
}

type ListMenuReply struct {
	List []*MenuReply `json:"list"`
}

func BuilderListMenuReply(menus []*model.Menu) *ListMenuReply {
	if menus == nil {
		return &ListMenuReply{List: []*MenuReply{}}
	}
	return &ListMenuReply{List: buildMenuList(menus, 0)}
}

func buildMenuList(menus []*model.Menu, parentID uint64) []*MenuReply {
	var list []*MenuReply
	for _, menu := range menus {
		if menu.ParentID == parentID {
			var roles []string
			if menu.Roles != "" {
				roles = strings.Split(menu.Roles, ",")
			}

			item := &MenuReply{
				ID:        int64(menu.ID),
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Meta: &MenuMeta{
					Title: menu.Title,
					Icon:  menu.Icon,
					Roles: roles,
				},
				Status:    int(menu.Status),
				UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
			children := buildMenuList(menus, menu.ID)
			if len(children) > 0 {
				item.Children = children
			}
			list = append(list, item)
		}
	}
	return list
}
