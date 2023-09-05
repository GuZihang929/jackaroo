package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"strconv"
	"user/model"

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

	res, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, in.Id)
	fmt.Println(err)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	return &rpc.UserInfoResponse{
		Id:       res.Id,
		Name:     res.Name.String,
		Age:      res.Age.Int64,
		Sex:      res.Sex.String,
		Phone:    res.Phone.String,
		Mail:     res.Mail.String,
		Identity: strconv.FormatInt(res.Identity.Int64, 10),
		Need:     res.Need.String,
	}, nil
}
