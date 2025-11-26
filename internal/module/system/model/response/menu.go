package response

type MenuTreeResp struct {
	ID        uint64          `json:"id"`
	Name      string          `json:"name"`
	Path      string          `json:"path"`
	Component string          `json:"component"`
	Redirect  string          `json:"redirect,omitempty"`
	Meta      MenuMeta        `json:"meta"`
	Children  []*MenuTreeResp `json:"children,omitempty"`
}

type MenuMeta struct {
	Title      string   `json:"title"`
	Icon       string   `json:"icon,omitempty"`
	IsHide     bool     `json:"isHide,omitempty"`
	IsHideTab  bool     `json:"isHideTab,omitempty"`
	Link       string   `json:"link,omitempty"`
	IsIframe   bool     `json:"isIframe,omitempty"`
	KeepAlive  bool     `json:"keepAlive,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	FixedTab   bool     `json:"fixedTab,omitempty"`
	ShowBadge  bool     `json:"showBadge,omitempty"`
	ActivePath string   `json:"activePath,omitempty"`
	IsFullPage bool     `json:"isFullPage,omitempty"`
}
