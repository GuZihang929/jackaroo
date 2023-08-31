package logic

import (
	"context"
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

func (l *MailRegisterLogic) MailSend(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.UserRpc.MailSendCode(l.ctx, &rpc.MailSendRequest{Phone: req.Account})
	if err != nil {
		return &types.RegisterResponse{
			Code: 500,
			Msg:  "出现异常",
		}, err
	}
	return &types.RegisterResponse{
		Code: 200,
		Msg:  "",
	}, nil
}

func (l *MailRegisterLogic) MailRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {

	_, err = l.svcCtx.UserRpc.Register(l.ctx, &rpc.RegisterRequest{
		Phone:    req.Account,
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
