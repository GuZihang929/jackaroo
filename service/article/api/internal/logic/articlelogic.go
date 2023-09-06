package logic

import (
	"article/api/internal/svc"
	"article/api/internal/types"
	"article/rpc/rpc"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Article 搜索文章
func (l *ArticleLogic) Article(req *types.ArtRequest) (resp *types.ArtResponse, err error) {
	fmt.Println(22222222222222)
	articles, err := l.svcCtx.ArticleRpc.Article(l.ctx, &rpc.ArtRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}
	resp = &types.ArtResponse{
		Code: articles.Code,
		Msg:  articles.Msq,
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
	return
}
