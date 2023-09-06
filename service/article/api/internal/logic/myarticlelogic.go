package logic

import (
	"article/api/internal/svc"
	"article/api/internal/types"
	"article/rpc/rpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type MyarticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyarticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyarticleLogic {
	return &MyarticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyarticleLogic) Myarticle() (resp *types.MyArtResponse, err error) {
	//i, _ := l.ctx.Value("userId").(json.Number).Int64()
	i := 499
	res, err := l.svcCtx.ArticleRpc.MyArticle(l.ctx, &rpc.MyArtRequest{
		ArtId: int64(i),
	})

	resp = &types.MyArtResponse{
		Code: res.Code,
		Msg:  res.Msq,
	}

	for _, date := range res.Date {
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
