package logic

import (
	"backstage/rpc/rpc"
	"context"

	"backstage/api/internal/svc"
	"backstage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetseajobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetseajobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetseajobLogic {
	return &GetseajobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetseajobLogic) Getseajob() (resp *types.JobResponse, err error) {

	res, err := l.svcCtx.BackstageRpc.GetSeaJob(l.ctx, &rpc.JobRequest{})

	if err != nil {
		return nil, err
	}
	resp = &types.JobResponse{
		Code:   res.Code,
		Msg:    res.Msg,
		Job:    res.Job,
		People: res.People,
	}
	return
}
