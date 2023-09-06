package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"job/rpc/internal/svc"
	"job/rpc/rpc"
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

func (l *LoginLogic) Login(in *rpc.SerchJobRequest) (*rpc.SerchJobResponse, error) {

	jobs := &[]Job{}

	err := l.svcCtx.DB.QueryRows(jobs, in.Sql)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@", err)
	if err != nil {
		return nil, err
	}

	resp := &rpc.SerchJobResponse{
		Code: 200,
		Msg:  "",
		Date: []*rpc.Job{},
	}
	//fmt.Println(jobs)

	for _, date := range *jobs {

		resp.Date = append(resp.Date, &rpc.Job{
			Id:          date.Id,
			CId:         date.CId,
			Company:     date.Company,
			Title:       date.Title,
			JobCategory: date.JobCategory,
			JobTypeName: date.JobTypeName,
			JobDetail:   date.JobDetail,
			JobLocation: date.JobLocation,
			PushTime:    date.PushTime,
			FetchTime:   date.FetchTime,
		})
	}

	return resp, err
}

type Job struct {
	Id          int64  `gorm:"uuid"`
	CId         string `gorm:"id"`
	Company     string `gorm:"company"`
	Title       string `gorm:"title"`
	JobCategory string `gorm:"job_category"`
	JobTypeName string `gorm:"job_type_name"`
	JobDetail   string `gorm:"job_detail"`
	JobLocation string `gorm:"job_location"`
	PushTime    string `gorm:"push_time"`
	FetchTime   string `gorm:"fetch_time"`
}
