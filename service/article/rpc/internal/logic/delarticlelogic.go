package logic

import (
	"context"

	"article/rpc/internal/svc"
	"article/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelArticleLogic {
	return &DelArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelArticleLogic) DelArticle(in *rpc.DelArtRequest) (*rpc.DelArtResponse, error) {

	err := l.svcCtx.ArtModel.Delete(l.ctx, in.Art)
	if err != nil {
		return &rpc.DelArtResponse{
			Code: 200,
			Msq:  "删除失败",
		}, err
	}

	return &rpc.DelArtResponse{
		Code: 200,
		Msq:  "删除成功",
	}, err
}
