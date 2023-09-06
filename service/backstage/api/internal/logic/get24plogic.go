package logic

import (
	"backstage/rpc/rpc"
	"context"
	"fmt"

	"backstage/api/internal/svc"
	"backstage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get24pLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet24pLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get24pLogic {
	return &Get24pLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get24pLogic) Get24p() (resp *types.PeoResponse, err error) {
	fmt.Println("------------")
	res, err := l.svcCtx.BackstageRpc.Get24P(l.ctx, &rpc.PeoRequest{})
	if err != nil {
		return nil, err
	}
	resp = &types.PeoResponse{
		Code:   res.Code,
		Msg:    res.Msg,
		Time:   res.Time,
		People: res.People,
	}

	return
}
