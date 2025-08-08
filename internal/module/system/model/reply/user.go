package reply

import (
	"server/internal/module/system/model"
	"strings"
)

type UserDetailReply struct {
	ID          int64    `json:"id"`
	Username    string   `json:"username"`
	Nickname    string   `json:"nickname"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Avatar      string   `json:"avatar"`
	Gender      int64    `json:"gender"`
	Status      int64    `json:"status"`
	IsAdmin     int64    `json:"isAdmin"`
	Province    string   `json:"province"`
	City        string   `json:"city"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Position    string   `json:"position"`
	Department  string   `json:"department"`
	JobTitle    string   `json:"jobTitle"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"createdAt"` // 注册时间
	LastLoginAt string   `json:"lastLoginAt,omitempty"`
	LastLoginIP string   `json:"lastLoginIP,omitempty"`

	Roles []string `json:"roles"`
}

func BuilderUserDetailReply(user *model.User) *UserDetailReply {
	roles := make([]string, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, r.Key)
	}

	return &UserDetailReply{
		ID:          int64(user.ID),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Phone:       user.Phone,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		Status:      user.Status,
		IsAdmin:     user.IsAdmin,
		Province:    user.Province,
		City:        user.City,
		District:    user.District,
		Address:     user.Address,
		Position:    user.Position,
		Department:  user.Department,
		JobTitle:    user.JobTitle,
		Tags:        strings.Split(user.Tags, ","),
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		LastLoginAt: user.LastLoginAt.Format("2006-01-02 15:04:05"),
		LastLoginIP: user.LastLoginIP,
		Roles:       roles,
	}
}
