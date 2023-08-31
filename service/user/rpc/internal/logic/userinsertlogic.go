package logic

import (
	"context"
	"fmt"
	"user/rpc/internal/svc"
	"user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInsertLogic {
	return &UserInsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInsertLogic) UserInsert(in *rpc.UserInsertRequest) (*rpc.UserInsertResponse, error) {
	//atoi, _ := strconv.Atoi(in.Identity)

	res, err := l.svcCtx.DB.ExecCtx(l.ctx, "update user set name=?,sex=?,age=?,identity=?,need=? where uuid = ?",
		in.Name, in.Sex, in.Age, in.Identity, in.Need, in.Uuid)

	if err != nil {
		return nil, err
	}

	fmt.Println(res)
	return &rpc.UserInsertResponse{
		Msg:  "出入成功",
		Code: 200,
	}, nil
}
