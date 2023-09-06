package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"jackaroo/service/position/model/model"
	"jackaroo/service/position/rpc/internal/svc"
	"jackaroo/service/position/rpc/types/position"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *position.UpdateRequest) (*position.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PositionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "公司不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	if in.Title != "" {
		res.Title = in.Title

	}
	if in.JobCategory != "" {
		res.JobCategory = in.JobCategory

	}
	if in.JobTypeName != "" {
		res.JobTypeName = in.JobTypeName

	}

	//// 声明一个 sql.NullString 类型的变量
	//var jobDetail sql.NullString
	//
	//// 检查 in.JobDetail 是否为空
	//if in.JobDetail == "" {
	//	// 如果 in.JobDetail 为空，则将 jobDetail 字段设置为无效（Valid = false）
	//	jobDetail.Valid = false
	//} else {
	//	// 如果 in.JobDetail 不为空，则将 jobDetail 字段设置为有效，并存储 in.JobDetail 的值
	//	jobDetail.Valid = true
	//	jobDetail.String = in.JobDetail
	//}
	//
	//// 声明一个 sql.NullString 类型的变量
	//var jobLocation sql.NullString
	//
	//// 检查 in.JobLocation 是否为空
	//if in.JobLocation == "" {
	//	// 如果 in.JobLocation 为空，则将 jobLocation 字段设置为无效（Valid = false）
	//	jobLocation.Valid = false
	//} else {
	//	// 如果 in.JobLocation 不为空，则将 jobLocation 字段设置为有效，并存储 in.JobLocation 的值
	//	jobLocation.Valid = true
	//	jobLocation.String = in.JobLocation
	//}

	if in.JobDetail != "" {
		res.JobDetail = in.JobDetail
	}

	if in.JobLocation != "" {
		res.JobLocation = in.JobLocation
	}
	err = l.svcCtx.PositionModel.Update(l.ctx, res)

	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &position.UpdateResponse{}, nil
}
