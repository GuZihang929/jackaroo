package logic

import (
	"context"
	"strconv"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailRegisterLogic {
	return &MailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailRegisterLogic) MailRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	atoi, err := strconv.Atoi(req.Code)
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &rpc.RegisterRequest{
		Phone:    req.Account,
		Code:     int64(atoi),
		Password: req.Password,
	})
	if err != nil {
		resp.Code = 500
		resp.Msg = "注册出现问题"
	}
	resp.Code = 200
	resp.Msg = "完成"
	return
}
