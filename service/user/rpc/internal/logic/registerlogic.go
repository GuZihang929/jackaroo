package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jackaroo/common"
	"sync"
	"time"
	"user/model"
	"user/rpc/internal/svc"
	"user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	// 判断手机号是否已经注册
	user, _ := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	sf := &common.Snowflake{
		Mutex:        sync.Mutex{},
		Timestamp:    time.Now().Unix(),
		Workerid:     124,
		Datacenterid: 1,
		Sequence:     01231,
	}
	uuid := sf.NextVal()
	if user == nil {
		fmt.Println("-=-=-=-=-=-=")
		u := &model.User{
			Uuid:     uuid,
			Phone:    sql.NullString{String: in.Phone, Valid: true},
			Password: sql.NullString{String: in.Password, Valid: true},
		}
		insert, e := l.svcCtx.UserModel.Insert(l.ctx, u)
		if e != nil {
			return nil, e
		}
		fmt.Println(insert)

		return &rpc.RegisterResponse{
			Name: "注册成功",
			Ok:   "200",
		}, nil
	}
	return &rpc.RegisterResponse{}, errors.New("已经存在")
}
