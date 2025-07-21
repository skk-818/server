package request

type PageInfo struct {
	Page     *int64 `json:"page" form:"page" validate:"required,notzero"`         // 页码
	PageSize *int64 `json:"pageSize" form:"pageSize" validate:"required,notzero"` // 每页大小
}

func (p *PageInfo) BuilderOffsetAndLimit() (offset, limit int) {
	offset = int((*p.Page - 1) * *p.PageSize)
	limit = int(*p.PageSize)
	return
}
