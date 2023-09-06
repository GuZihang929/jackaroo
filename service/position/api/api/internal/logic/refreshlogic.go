package logic

import (
	"context"
	"jackaroo/service/position/rpc/types/position"

	"jackaroo/service/position/api/api/internal/svc"
	"jackaroo/service/position/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshRequest) (resp *types.RefreshResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.PositionRpc.Refresh(l.ctx, &position.RefreshRequest{
		Cid: req.Cid,
	})
	if err != nil {
		return nil, err
	}
	return
}
