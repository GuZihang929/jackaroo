package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"jackaroo/service/company/model/model"

	"jackaroo/service/company/rpc/internal/svc"
	"jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *company.RemoveRequest) (*company.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "公司不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.CompanyModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &company.RemoveResponse{}, nil
}
