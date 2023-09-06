package logic

import (
	"context"
	"jackaroo/service/company/rpc/types/company"

	"jackaroo/service/company/api/api/internal/svc"
	"jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyRpc.Detail(l.ctx, &company.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:      res.Id,
		Company: res.Company,
		Number:  res.Number,
		Brief:   res.Brief,
		Picture: res.Picture,
	}, nil
}
