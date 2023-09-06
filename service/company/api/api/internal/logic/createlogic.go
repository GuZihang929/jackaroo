package logic

import (
	"context"
	"jackaroo/service/company/rpc/types/company"

	"jackaroo/service/company/api/api/internal/svc"
	"jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyRpc.Create(l.ctx, &company.CreateRequest{
		Company: req.Company,
		Brief:   req.Brief,
		Picture: req.Picture,
	})

	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
