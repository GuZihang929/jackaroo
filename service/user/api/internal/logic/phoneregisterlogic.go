package logic

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRegisterLogic {
	return &PhoneRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneRegisterLogic) PhoneRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {

	return
}
