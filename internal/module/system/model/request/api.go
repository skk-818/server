package request

type CreateApiReq struct {
	Name        string `json:"name" validate:"required"`
	Path        string `json:"path" validate:"required"`
	Method      string `json:"method" validate:"required"`
	Description string `json:"description" validate:"required"`
	Group       string `json:"group" validate:"required"`
}

type DeleteApiReq struct {
	Id int64 `json:"id" validate:"required,min=1"`
}

type UpdateApiReq struct {
	Id          int64  `json:"id" validate:"required,min=1"`
	Name        string `json:"name" validate:"required"`
	Path        string `json:"path" validate:"required"`
	Method      string `json:"method" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      *int64 `json:"status" validate:"required"`
	Group       string `json:"group" validate:"required"`
}

type ApiDetailReq struct {
	Id int64 `json:"id" validate:"required,min=1"`
}

type ApiListReq struct {
	PageInfo
	Path   string `json:"path"`
	Method string `json:"method"`
	Status int64  `json:"status"`
}
