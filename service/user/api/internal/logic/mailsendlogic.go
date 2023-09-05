package logic

import (
	"context"
	"user/rpc/rpc"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailSendLogic {
	return &MailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailSendLogic) MailSend(req *types.MailSendRequest) (resp *types.MailSendResponse, err error) {
	_, err = l.svcCtx.UserRpc.MailSendCode(l.ctx, &rpc.MailSendRequest{Phone: req.Account})
	if err != nil {
		return &types.MailSendResponse{
			Code: 500,
			Msg:  "出现异常",
		}, err
	}
	return &types.MailSendResponse{
		Code: 200,
		Msg:  "",
	}, nil
	return
}
