package logic

import (
	"article/rpc/rpc"
	"context"

	"article/api/internal/svc"
	"article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelarticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelarticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelarticleLogic {
	return &DelarticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelarticleLogic) Delarticle(req *types.DelArtRequest) (resp *types.DelArtResponse, err error) {

	article, err := l.svcCtx.ArticleRpc.DelArticle(l.ctx, &rpc.DelArtRequest{Art: req.ArtId})
	if err != nil {
		return &types.DelArtResponse{
			Code: 200,
			Msg:  "删除失败",
		}, err
	}

	resp = &types.DelArtResponse{
		Code: article.Code,
		Msg:  article.Msq,
	}
	return
}
