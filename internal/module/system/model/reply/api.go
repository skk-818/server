package reply

import (
	"server/internal/module/system/model"
)

type ApiDetailReply struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func BuilderApiDetailReply(api *model.Api) *ApiDetailReply {
	return &ApiDetailReply{
		Id:          int64(api.ID),
		Name:        api.Name,
		Path:        api.Path,
		Method:      api.Method,
		Description: api.Description,
		Group:       api.Group,
		Status:      api.Status,
		CreatedAt:   api.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   api.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type ApiListReply struct {
	PageReply
	List []*ApiReply `json:"list"`
}

type ApiReply struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
