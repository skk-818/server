package request

type CreateApiReq struct {
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
