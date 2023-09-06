package logic

import (
	"context"
	"jackaroo/service/company/rpc/types/company"

	"jackaroo/service/company/api/api/internal/svc"
	"jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshRequest) (resp *types.RefreshResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyRpc.Refresh(l.ctx, &company.RefreshRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.RefreshResponse{
		Id:      req.Id,
		Company: res.Company,
		Number:  res.Number,
		Brief:   res.Brief,
		Picture: res.Picture,
	}, nil
}
