package request

type CreateApiReq struct {
	Name        string `json:"name" validate:"required"`
	Path        string `json:"path" validate:"required"`
	Method      string `json:"method" validate:"required"`
	Description string `json:"description" validate:"required"`
	Group       string `json:"group" validate:"required"`
}

type DeleteApiReq struct {
	ID int64 `json:"id" validate:"required"`
}

type UpdateApiReq struct {
	ID          int64  `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Path        string `json:"path" validate:"required"`
	Method      string `json:"method" validate:"required"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Status      int64  `json:"status"`
}

type GetApiReq struct {
	ID int64 `json:"id" validate:"required"`
}

type ApiListReq struct {
	PageInfo
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
	Status int64  `json:"status" form:"status"`
}
