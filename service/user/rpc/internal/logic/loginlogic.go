package logic

import (
	"context"
	"fmt"
	"jackaroo/common"
	"strconv"
	"user/model"
	"user/rpc/internal/svc"
	"user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *rpc.LoginRequest) (*rpc.LoginResponse, error) {

	var user model.User
	err := l.svcCtx.DB.QueryRowCtx(l.ctx, &user, "select * from user where mail = ?", in.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	value := user.Auth.Int64
	fmt.Println(value)
	Code := 0
	if value == 1 {
		Code = 1
	}
	fmt.Println(Code)

	token, _ := common.GenToken(int(user.Uuid))
	fmt.Println(token)
	return &rpc.LoginResponse{
		Name: token,
		Ok:   strconv.Itoa(Code),
	}, nil
}
