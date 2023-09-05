package logic

import (
	"context"
	"xiangxiang/jackaroo/service/position/rpc/types/position"

	"xiangxiang/jackaroo/service/position/api/api/internal/svc"
	"xiangxiang/jackaroo/service/position/api/api/internal/types"

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
	res, err := l.svcCtx.PositionRpc.Create(l.ctx, &position.CreateRequest{
		Cid:         req.Cid,
		Title:       req.Title,
		JobCategory: req.JobCategory,
		JobTypeName: req.JobTypeName,
		JobDetail:   req.JobDetail,
		JobLocation: req.JobLocation,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
