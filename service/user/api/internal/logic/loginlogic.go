package logic

import (
	"context"
	"fmt"
	"strconv"
	"user/rpc/rpc"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {

	loginResponse, err := l.svcCtx.UserRpc.Login(l.ctx, &rpc.LoginRequest{
		Name:     req.Account,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResponse{
			AccessToken:  "",
			AccessExpire: 0,
		}, err
	}
	fmt.Println(loginResponse.Name)
	fmt.Println(loginResponse.Ok)
	atoi, _ := strconv.Atoi(loginResponse.Ok)
	return &types.LoginResponse{
		Code:         atoi,
		AccessToken:  loginResponse.Name,
		AccessExpire: 60 * 60,
	}, err
}
