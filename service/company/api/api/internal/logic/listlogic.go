package logic

import (
	"context"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"xiangxiang/jackaroo/service/company/api/api/internal/svc"
	"xiangxiang/jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp []*types.ListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyRpc.List(l.ctx, &company.ListRequest{})
	if err != nil {
		return nil, err
	}
	companyList := make([]*types.ListResponse, 0)
	for _, item := range res.Data {
		companyList = append(companyList, &types.ListResponse{
			Id:      item.Id,
			Company: item.Company,
			Number:  item.Number,
			Brief:   item.Brief,
			Picture: item.Picture,
		})

	}
	return companyList, nil
}
