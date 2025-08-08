package usecase

import (
	"context"
	"go.uber.org/zap"
	"server/internal/core/logger"
	"server/internal/module/system/model/reply"
	"server/internal/module/system/model/request"
	"server/internal/module/system/usecase/repo"
	"server/pkg/errorx"
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

func (u *ApiUsecase) Create(ctx context.Context, params *request.CreateApiReq) error {
	return nil
}

func (u *ApiUsecase) Delete(ctx context.Context, params *request.DeleteApiReq) error {
	return nil
}

func (u *ApiUsecase) Update(ctx context.Context, params *request.UpdateApiReq) error {
	return nil
}

func (u *ApiUsecase) Detail(ctx context.Context, params *request.ApiDetailReq) (*reply.ApiDetailReply, error) {
	api, err := u.apiRepo.Find(ctx, params.Id)
	if err != nil {
		u.logger.Error("[ApiUsecase] apiRepo.Find error", zap.Error(err))
		return nil, err
	}
	if api == nil {
		u.logger.Error("[ApiUsecase] api not found", zap.Error(err))
		return nil, errorx.ErrApiNotFound
	}

	apiReply := &reply.ApiDetailReply{
		Id:          int64(api.ID),
		Name:        api.Name,
		Path:        api.Path,
		Method:      api.Method,
		Description: api.Description,
		Group:       api.Group,
		Status:      api.Status,
		CreatedAt:   api.CreatedAt,
		UpdatedAt:   api.UpdatedAt,
	}

	return apiReply, nil
}

func (u *ApiUsecase) List(ctx context.Context, params *request.ApiListReq) (*reply.ApiListReply, error) {
	apis, total, err := u.apiRepo.List(ctx, params)
	if err != nil {
		u.logger.Error("[ApiUsecase.List] apiRepo.List Err", zap.Error(err))
		return nil, err
	}

	var apiReplyList []*reply.ApiReply
	for _, api := range apis {
		apiReplyList = append(apiReplyList, &reply.ApiReply{
			Id:          int64(api.ID),
			Name:        api.Name,
			Path:        api.Path,
			Method:      api.Method,
			Description: api.Description,
			Group:       api.Group,
			Status:      api.Status,
			CreatedAt:   api.CreatedAt,
			UpdatedAt:   api.UpdatedAt,
		})
	}

	return &reply.ApiListReply{
		PageReply: reply.PageReply{
			Total:    total,
			Page:     *params.Page,
			PageSize: *params.PageSize,
		},
		List: apiReplyList,
	}, nil
}
