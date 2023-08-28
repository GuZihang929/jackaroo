package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *rpc.UserInfoRequest) (*rpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.UserInfoResponse{}, nil
}
