package logic

import (
	"context"

	"job/api/internal/svc"
	"job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobLogic {
	return &JobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobLogic) Job(req *types.JobRequest) (resp *types.JobResponse, err error) {

	var job types.Job
	err = l.svcCtx.DB.QueryRow(&job, "select * from job where uuid = ?", req.Id)
	if err != nil {
		resp = &types.JobResponse{
			Code: 202,
			Msg:  "获取失败",
			Date: types.Job{},
		}
	}
	resp = &types.JobResponse{
		Code: 200,
		Msg:  "获取成功",
		Date: job,
	}
	return
}
