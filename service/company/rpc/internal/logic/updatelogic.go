package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"jackaroo/service/company/model/model"

	"github.com/zeromicro/go-zero/core/logx"
	"jackaroo/service/company/rpc/internal/svc"
	"jackaroo/service/company/rpc/types/company"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *company.UpdateRequest) (*company.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Company != "" {
		res.Company = in.Company
	}
	if in.Brief != "" {
		res.Brief = in.Brief
	}
	if in.Picture != "" {
		res.Picture = in.Picture
	}

	err = l.svcCtx.CompanyModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &company.UpdateResponse{}, nil
}
