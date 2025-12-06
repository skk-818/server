package reply

import "server/internal/module/system/model"

type GetRoleReply struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func BuilderGetRoleReply(role *model.Role) *GetRoleReply {
	if role == nil {
		return nil
	}

	return &GetRoleReply{
		ID:        int64(role.ID),
		Name:      role.Name,
		Status:    int(role.Status),
		Remark:    role.Remark,
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type RoleReply struct {
	ID          int64  `json:"roleId"`
	Name        string `json:"roleName"`
	Key         string `json:"roleCode"`
	Status      int    `json:"enabled"`
	Description string `json:"description"`
	CreatedAt   string `json:"createTime"`
}

type ListRoleReply struct {
	List     []*RoleReply `json:"list"`
	Total    int64        `json:"total"`
	Page     int64        `json:"page"`
	PageSize int64        `json:"pageSize"`
}

func BuilderListRoleReply(roles []*model.Role, total int64, page, pageSize int64) *ListRoleReply {
	if roles == nil {
		return &ListRoleReply{
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		}
	}
	var list []*RoleReply
	for _, role := range roles {
		list = append(list, &RoleReply{
			ID:          int64(role.ID),
			Name:        role.Name,
			Key:         role.Key,
			Status:      int(role.Status),
			Description: role.Remark,
			CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &ListRoleReply{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}
