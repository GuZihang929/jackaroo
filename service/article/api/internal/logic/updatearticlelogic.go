package logic

import (
	"article/rpc/rpc"
	"context"

	"article/api/internal/svc"
	"article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatearticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatearticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatearticleLogic {
	return &UpdatearticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatearticleLogic) Updatearticle(req *types.UpdateArtRequest) (resp *types.UpdateArtResponse, err error) {

	article, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &rpc.UpdateArtRequest{
		ArtId: req.ArtId,
		Name:  req.Name,
		Info:  req.Info,
		Tag:   req.Tag,
		Time:  req.Time,
	})
	if err != nil {
		return &types.UpdateArtResponse{
			Code: 200,
			Msg:  "修改失败",
		}, err
	}

	resp = &types.UpdateArtResponse{
		Code: article.Code,
		Msg:  article.Msq,
	}

	return
}
