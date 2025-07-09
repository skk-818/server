package reply

type PageReply struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	List     any   `json:"list"`
}
