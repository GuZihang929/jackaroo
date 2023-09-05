package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"xiangxiang/jackaroo/service/company/model/model"

	"xiangxiang/jackaroo/service/company/rpc/internal/svc"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *company.ListRequest) (*company.ListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.CompanyModel.FindAll(l.ctx)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "公司不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	companyList := make([]*company.DetailResponse, 0)
	for _, item := range list {
		companyList = append(companyList, &company.DetailResponse{
			Id:      item.Id,
			Company: item.Company,
			Number:  item.Number,
			Brief:   item.Brief,
			Picture: item.Picture,
		})
	}
	return &company.ListResponse{
		Data: companyList,
	}, nil
}
