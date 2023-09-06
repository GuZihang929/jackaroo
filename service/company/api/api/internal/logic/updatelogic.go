package logic

import (
	"context"
	"jackaroo/service/company/rpc/types/company"

	"jackaroo/service/company/api/api/internal/svc"
	"jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CompanyRpc.Update(l.ctx, &company.UpdateRequest{
		Id:      req.Id,
		Company: req.Company,
		Brief:   req.Brief,
		Picture: req.Picture,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
