package logic

import (
	"context"
	"fmt"
	"strconv"

	"backstage/rpc/internal/svc"
	"backstage/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get24pLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGet24pLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get24pLogic {
	return &Get24pLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Get24pLogic) Get24P(in *rpc.PeoRequest) (*rpc.PeoResponse, error) {

	m, err := l.svcCtx.Redis.Hgetall("User_Sum")
	if err != nil {
		fmt.Println("===============")
		return nil, err
	}

	resq := &rpc.PeoResponse{
		Code:   200,
		Msg:    "获取成功",
		Time:   []int64{},
		People: []int64{},
	}
	fmt.Println(m)

	for i := 0; i < 24; i++ {

		atoi, _ := strconv.Atoi(m[strconv.Itoa(i)])
		resq.Time = append(resq.Time, int64(i))
		resq.People = append(resq.People, int64(atoi))
	}

	return resq, err
}
