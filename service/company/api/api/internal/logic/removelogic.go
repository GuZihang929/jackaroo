package logic

import (
	"context"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"xiangxiang/jackaroo/service/company/api/api/internal/svc"
	"xiangxiang/jackaroo/service/company/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CompanyRpc.Remove(l.ctx, &company.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
