package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"jackaroo/service/company/rpc/types/company"
	"jackaroo/service/position/model/model"
	"jackaroo/service/position/rpc/internal/svc"
	"jackaroo/service/position/rpc/types/position"
	"time"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *position.CreateRequest) (*position.CreateResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(l.ctx.Deadline())
	fmt.Println("---------------------------------------------------------------")
	companyRes, err := l.svcCtx.CompanyRpc.Detail(l.ctx, &company.DetailRequest{
		Id: in.Cid,
	})
	fmt.Println("Cid:", in.Cid)
	if err != nil {
		return nil, err
	}
	if companyRes.Id < 0 {
		return nil, status.Error(500, "企业不存在")
	}

	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")

	newPosition := model.Position{
		Cid:         companyRes.Id,
		Company:     companyRes.Company,
		Title:       in.Title,
		JobCategory: in.JobCategory,
		JobTypeName: in.JobTypeName,
		JobDetail:   in.JobDetail,
		JobLocation: in.JobLocation,
		PushTime:    currentTimeString,
		FetchTime:   currentTimeString,
	}
	res, err := l.svcCtx.PositionModel.Insert(l.ctx, &newPosition)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newPosition.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &position.CreateResponse{
		Id: newPosition.Id,
	}, nil
}
