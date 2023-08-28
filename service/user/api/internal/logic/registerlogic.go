package logic

import (
	"context"
	"user/rpc/user"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Phone:    req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Name: res.Name,
		Ok:   res.Ok,
	}, nil
}
