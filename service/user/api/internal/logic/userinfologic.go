package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"user/rpc/rpc"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {

	uid, err := l.ctx.Value("userId").(json.Number).Int64()

	if err != nil {
		fmt.Println(err)
	}

	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &rpc.UserInfoRequest{
		Id: uid,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:       info.Id,
		Name:     info.Name,
		Sex:      info.Sex,
		Age:      int(info.Age),
		Phone:    info.Phone,
		Mail:     info.Mail,
		Identity: info.Identity,
		Need:     info.Need,
	}, nil

	return
}
