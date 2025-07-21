package request

type CreateApiReq struct {
	Name        string `json:"name" validate:"required"`
	Path        string `json:"path" validate:"required"`
	Method      string `json:"method" validate:"required"`
	Description string `json:"description" validate:"required"`
	Group       string `json:"group" validate:"required"`
}

type DeleteApiReq struct {
}

type UpdateApiReq struct {
}

type GetApiReq struct {
}

type ApiListReq struct {
	PageInfo
	Path   string `json:"path"`
	Method string `json:"method"`
	Status int64  `json:"status"`
}
