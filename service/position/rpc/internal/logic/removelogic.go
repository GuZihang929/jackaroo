package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"jackaroo/service/position/model/model"

	"jackaroo/service/position/rpc/internal/svc"
	"jackaroo/service/position/rpc/types/position"

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

func (l *RemoveLogic) Remove(in *position.RemoveRequest) (*position.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PositionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "职位不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	err = l.svcCtx.PositionModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &position.RemoveResponse{}, nil
}
