package logic

import (
	"context"
	"fmt"

	"job/api/internal/svc"
	"job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagLogic {
	return &TagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagLogic) Tag(req *types.TagAddRequest) (resp *types.TagAddResponse, err error) {
	fmt.Println("-----------------")
	_, err = l.svcCtx.Redis.Hincrby("Job", req.Tag, 1)
	if err != nil {
		resp = &types.TagAddResponse{
			Code: 202,
			Msg:  "获取失败",
			Date: "",
		}
	}
	resp = &types.TagAddResponse{
		Code: 200,
		Msg:  "添加成功",
		Date: "",
	}
	return
}
