package logic

import (
	"article/api/internal/svc"
	"article/api/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeRequest) (resp *types.LikeResponse, err error) {

	_, err = l.svcCtx.DB.Exec("update article set like_sum = like_sum +? where uuid = ?", req.Sum, req.ArtId)
	if err != nil {
		fmt.Println(err)
	}
	resp = &types.LikeResponse{
		Code: 200,
		Msg:  "添加成功",
		Date: types.Articles{},
	}

	return
}
