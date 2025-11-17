package biz

import (
	"context"
	"server/internal/core/logger"
	"server/internal/module/system/biz/repo"
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
	return nil
}

func (u ApiUsecase) Delete(ctx context.Context, req *request.DeleteApiReq) error {
	return nil
}

func (u ApiUsecase) Update(ctx context.Context, req *request.UpdateApiReq) error {
	return nil
}

func (u ApiUsecase) Get(ctx context.Context, req *request.GetApiReq) (*reply.GetApiReply, error) {
	return nil, nil
}

func (u ApiUsecase) List(ctx context.Context, req *request.ApiListReq) ([]*reply.ListApiReply, error) {
	return nil, nil
}
