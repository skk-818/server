package reply

import "time"

type ApiDetailReply struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Group       string    `json:"group"`
	Status      int64     `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ApiListReply struct {
	PageReply
	List []*ApiReply `json:"list"`
}

type ApiReply struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Group       string    `json:"group"`
	Status      int64     `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
