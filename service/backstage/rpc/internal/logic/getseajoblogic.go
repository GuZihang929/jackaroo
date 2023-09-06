package logic

import (
	"context"
	"fmt"
	"strconv"

	"backstage/rpc/internal/svc"
	"backstage/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeaJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSeaJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeaJobLogic {
	return &GetSeaJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSeaJobLogic) GetSeaJob(in *rpc.JobRequest) (*rpc.JobResponse, error) {

	m, err := l.svcCtx.Redis.Hgetall("Job")
	if err != nil {
		return nil, err
	}
	resq := &rpc.JobResponse{
		Code:   200,
		Msg:    "获取成功",
		Job:    []string{},
		People: []int64{},
	}
	fmt.Println(m)

	for k, v := range m {
		resq.Job = append(resq.Job, k)
		atoi, _ := strconv.Atoi(v)
		resq.People = append(resq.People, int64(atoi))
	}

	return resq, err
}
