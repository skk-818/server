package reply

import "server/internal/module/system/model"

type GetApiReply struct {
	*model.Api
}

type ApiReply struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ListApiReply struct {
	List     []*ApiReply `json:"list"`
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
}

func BuilderListApiReply(apis []*model.Api, total int64, page, pageSize int64) *ListApiReply {
	if apis == nil {
		return &ListApiReply{
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		}
	}
	var list []*ApiReply
	for _, api := range apis {
		list = append(list, &ApiReply{
			ID:          int64(api.ID),
			Name:        api.Name,
			Path:        api.Path,
			Method:      api.Method,
			Description: api.Description,
			Group:       api.Group,
			Status:      int(api.Status),
			CreatedAt:   api.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   api.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &ListApiReply{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}
