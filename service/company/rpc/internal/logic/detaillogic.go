package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"xiangxiang/jackaroo/service/company/model/model"

	"xiangxiang/jackaroo/service/company/rpc/internal/svc"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *company.DetailRequest) (*company.DetailResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "企业不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &company.DetailResponse{
		Id:      res.Id,
		Company: res.Company,
		Number:  res.Number,
		Brief:   res.Brief,
		Picture: res.Picture,
	}, nil
}
