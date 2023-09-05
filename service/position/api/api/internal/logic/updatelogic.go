package logic

import (
	"context"
	"xiangxiang/jackaroo/service/position/rpc/types/position"

	"xiangxiang/jackaroo/service/position/api/api/internal/svc"
	"xiangxiang/jackaroo/service/position/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.PositionRpc.Update(l.ctx, &position.UpdateRequest{
		Id:          req.Id,
		Title:       req.Title,
		JobCategory: req.JobCategory,
		JobTypeName: req.JobTypeName,
		JobDetail:   req.JobDetail,
		JobLocation: req.JobLocation,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
