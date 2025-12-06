package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
	"server/internal/module/system/model"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
)

type ApiUsecase struct {
	logger  logger.Logger
	apiRepo repo.ApiRepo
}

func NewApiUsecase(
	logger logger.Logger,
	apiRepo repo.ApiRepo,
) *ApiUsecase {
	return &ApiUsecase{
		logger:  logger,
		apiRepo: apiRepo,
	}
}

func (u ApiUsecase) Create(ctx context.Context, req *request.CreateApiReq) error {
	api := &model.Api{
		Name:        req.Name,
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Description,
		Group:       req.Group,
		Status:      1,
	}
	return u.apiRepo.Create(ctx, api)
}

func (u ApiUsecase) Delete(ctx context.Context, req *request.DeleteApiReq) error {
	return u.apiRepo.Delete(ctx, req.ID)
}

func (u ApiUsecase) Update(ctx context.Context, req *request.UpdateApiReq) error {
	api := &model.Api{
		BaseModel:   model.BaseModel{ID: uint64(req.ID)},
		Name:        req.Name,
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Description,
		Group:       req.Group,
		Status:      req.Status,
	}
	return u.apiRepo.Update(ctx, api)
}

func (u ApiUsecase) Get(ctx context.Context, req *request.GetApiReq) (*reply.GetApiReply, error) {
	api, err := u.apiRepo.Find(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &reply.GetApiReply{Api: api}, nil
}

func (u ApiUsecase) List(ctx context.Context, req *request.ApiListReq) (*reply.ListApiReply, error) {
	list, total, err := u.apiRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	page, pageSize := int64(1), int64(20)
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		pageSize = *req.PageSize
	}
	return reply.BuilderListApiReply(list, total, page, pageSize), nil
}
