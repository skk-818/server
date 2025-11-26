package request

type MenuListReq struct {
	PageInfo
	Name     string `json:"name"`
	ParentID uint64 `json:"parentId"`
}

type MenuCreateReq struct {
	ParentID   uint64 `json:"parentId"`
	Name       string `json:"name" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Path       string `json:"path" binding:"required"`
	Component  string `json:"component"`
	Icon       string `json:"icon"`
	Redirect   string `json:"redirect"`
	Link       string `json:"link"`
	Roles      string `json:"roles"`
	IsIframe   int64  `json:"isIframe"`
	Hidden     int64  `json:"hidden"`
	HideTab    int64  `json:"hideTab"`
	KeepAlive  int64  `json:"keepAlive"`
	FullPage   int64  `json:"fullPage"`
	FixedTab   int64  `json:"fixedTab"`
	ShowBadge  int64  `json:"showBadge"`
	TextBadge  string `json:"textBadge"`
	ActivePath string `json:"activePath"`
	Sort       int64  `json:"sort"`
	Status     int64  `json:"status"`
}

type MenuUpdateReq struct {
	ID uint64 `json:"id" binding:"required"`
	MenuCreateReq
}

type MenuDeleteReq struct {
	ID uint64 `json:"id" binding:"required"`
}
