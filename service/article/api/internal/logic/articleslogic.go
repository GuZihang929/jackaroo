package logic

import (
	"article/rpc/rpc"
	"context"
	"fmt"

	"article/api/internal/svc"
	"article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlesLogic) Articles(req *types.ArtsRequest) (*types.ArtsResponse, error) {
	articles, err := l.svcCtx.ArticleRpc.Articles(l.ctx, &rpc.ArtsRequest{
		NextID:   req.NextID,
		PageSize: req.PageSize,
	})
	if err != nil {
		fmt.Println("---------", err)
		return nil, err
	}

	resp := &types.ArtsResponse{
		Code: articles.Code,
		Date: []types.Articles{},
	}

	for _, date := range articles.Date {
		resp.Date = append(resp.Date, types.Articles{
			Uuid:    date.Uuid,
			UserId:  date.UserId,
			Name:    date.Name,
			Info:    date.Info,
			Tag:     date.Tag,
			Time:    date.Time,
			LikeSum: date.LikeSum,
		})
	}
	return resp, err
}
