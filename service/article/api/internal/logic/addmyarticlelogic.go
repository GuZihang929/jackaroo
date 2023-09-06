package logic

import (
	"article/rpc/rpc"
	"context"
	"encoding/json"
	"fmt"

	"article/api/internal/svc"
	"article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddmyarticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddmyarticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddmyarticleLogic {
	return &AddmyarticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddmyarticleLogic) Addmyarticle(req *types.AddArtRequest) (resp *types.AddArtResponse, err error) {

	fmt.Println(req)
	uid, err := l.ctx.Value("userId").(json.Number).Int64()

	article, err := l.svcCtx.ArticleRpc.AddMyArticle(l.ctx, &rpc.AddArtRequest{
		Name:   req.Name,
		Info:   req.Info,
		Tag:    req.Tag,
		Time:   req.Time,
		UserId: uid,
	})

	if err != nil {
		return &types.AddArtResponse{
			Code: 200,
			Msg:  "添加失败",
		}, err
	}

	resp = &types.AddArtResponse{
		Code: article.Code,
		Msg:  article.Msq,
	}

	return
}
