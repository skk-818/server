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
	Path    string `json:"path"`
	Methods string `json:"methods"`
	Status  int64  `json:"status"`
}
