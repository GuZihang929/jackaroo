package logic

import (
	"context"
	"database/sql"
	"fmt"
	"jackaroo/common"
	"strconv"
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

	get, err := l.svcCtx.Redis.Get(in.Phone)
	if err != nil {
		return nil, err
	}

	sf := &common.Snowflake{
		Mutex:        sync.Mutex{},
		Timestamp:    time.Now().Unix(),
		Workerid:     124,
		Datacenterid: 1,
		Sequence:     01231,
	}

	uuid := sf.NextVal()

	if get == strconv.FormatInt(in.Code, 10) {
		_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
			Uuid:     uuid,
			Password: sql.NullString{String: in.Password, Valid: true},
			Mail:     sql.NullString{String: in.Phone, Valid: true},
			Auth:     sql.NullInt64{Int64: 0, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("注册完成")

	return &rpc.RegisterResponse{}, nil
}
