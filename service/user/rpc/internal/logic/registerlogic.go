package logic

import (
	"context"
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
	fmt.Println("1111111111111111")
	get, err := l.svcCtx.Redis.Get(in.Phone)
	if err != nil {
		return nil, err
	}
	fmt.Println(in)

	sf := &common.Snowflake{
		Mutex:        sync.Mutex{},
		Timestamp:    time.Now().Unix(),
		Workerid:     124,
		Datacenterid: 1,
		Sequence:     01231,
	}

	uuid := sf.NextVal()
	var user model.User

	if get == strconv.FormatInt(in.Code, 10) {
		err = l.svcCtx.DB.QueryRow(&user, "select * from user where mail = ?", in.Phone)
		fmt.Println(user)
		fmt.Println("===================1")
		if user.Mail.String != "" {
			fmt.Println("===================2")
			_, err = l.svcCtx.DB.Exec("update user set  password = ? where mail = ?", in.Password, in.Phone)
			fmt.Println(err)
			fmt.Println("===================3")
			if err != nil {
				return &rpc.RegisterResponse{
					Name: "保存有误",
					Ok:   "400",
				}, err
			}
		} else {
			fmt.Println("===================4")
			_, err = l.svcCtx.DB.Exec("insert into user (uuid,mail,password) values (?,?,?)", uuid, in.Phone, in.Password)
			if err != nil {
				return &rpc.RegisterResponse{
					Name: "保存有误",
					Ok:   "400",
				}, err
			}
		}

	} else {

		return &rpc.RegisterResponse{
			Name: "验证码有误",
			Ok:   "400",
		}, err
	}

	return &rpc.RegisterResponse{
		Name: "注册成功",
		Ok:   "200",
	}, nil
}
