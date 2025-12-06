package request

type PageInfo struct {
	Page     *int64 `json:"page" form:"current"`
	PageSize *int64 `json:"pageSize" form:"size"`
}

func (p *PageInfo) BuilderOffsetAndLimit() (offset, limit int) {
	if p.Page == nil || p.PageSize == nil {
		return 0, 10
	}
	offset = int((*p.Page - 1) * *p.PageSize)
	limit = int(*p.PageSize)
	return
}
