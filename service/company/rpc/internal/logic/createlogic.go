package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"xiangxiang/jackaroo/service/company/model/model"

	"xiangxiang/jackaroo/service/company/rpc/internal/svc"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *company.CreateRequest) (*company.CreateResponse, error) {
	// todo: add your logic here and delete this line
	newCompany := model.Company{
		Company: in.Company,
		Brief:   in.Brief,
		Picture: in.Picture,
	}
	result, err := l.svcCtx.CompanyModel.Insert(l.ctx, &newCompany)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newCompany.Id, err = result.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &company.CreateResponse{
		Id: newCompany.Id,
	}, nil
}
