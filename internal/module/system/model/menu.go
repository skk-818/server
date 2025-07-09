package model

type Menu struct {
	BaseModel

	ParentID  uint64 `gorm:"not null;default:0" json:"parentId"`            // 父菜单 ID，顶级菜单为 0
	Name      string `gorm:"size:64;not null;unique" json:"name"`           // 路由名称（唯一标识）
	Title     string `gorm:"size:128;not null" json:"title"`                // 菜单标题
	Path      string `gorm:"size:255;not null" json:"path"`                 // 路由路径
	Component string `gorm:"size:255;not null;default:''" json:"component"` // 对应前端组件路径
	Icon      string `gorm:"size:128;default:''" json:"icon"`               // 图标
	Redirect  string `gorm:"size:255;default:''" json:"redirect"`           // 重定向路径
	Link      string `gorm:"size:255;default:''" json:"link"`               // iframe 或外链地址

	IsIframe   int64  `gorm:"not null;default:0" json:"isIframe"`    // 是否 iframe 链接
	Hidden     int64  `gorm:"not null;default:0" json:"hidden"`      // 是否隐藏菜单
	HideTab    int64  `gorm:"not null;default:0" json:"hideTab"`     // 是否隐藏标签页
	KeepAlive  int64  `gorm:"not null;default:1" json:"keepAlive"`   // 是否缓存
	FullPage   int64  `gorm:"not null;default:0" json:"fullPage"`    // 是否全屏显示
	FixedTab   int64  `gorm:"not null;default:0" json:"fixedTab"`    // 是否固定标签页
	ShowBadge  int64  `gorm:"not null;default:0" json:"showBadge"`   // 是否展示小圆点
	TextBadge  string `gorm:"size:64;default:''" json:"textBadge"`   // 显示文本徽章
	ActivePath string `gorm:"size:255;default:''" json:"activePath"` // 激活的路径
	Sort       int64  `gorm:"not null;default:0" json:"sort"`        // 排序字段

}

var MenuCol = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	ParentID   string
	Name       string
	Title      string
	Path       string
	Component  string
	Icon       string
	Redirect   string
	Link       string
	IsIframe   string
	Hidden     string
	HideTab    string
	KeepAlive  string
	FullPage   string
	FixedTab   string
	ShowBadge  string
	TextBadge  string
	ActivePath string
	Sort       string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	ParentID:   "parent_id",
	Name:       "name",
	Title:      "title",
	Path:       "path",
	Component:  "component",
	Icon:       "icon",
	Redirect:   "redirect",
	Link:       "link",
	IsIframe:   "is_iframe",
	Hidden:     "hidden",
	HideTab:    "hide_tab",
	KeepAlive:  "keep_alive",
	FullPage:   "full_page",
	FixedTab:   "fixed_tab",
	ShowBadge:  "show_badge",
	TextBadge:  "text_badge",
	ActivePath: "active_path",
	Sort:       "sort",
}
