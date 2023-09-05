package logic

import (
	"article/api/internal/svc"
	"article/api/internal/types"
	"article/model"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArtLogic {
	return &ArtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtLogic) Art(req *types.ARequest) (resp *types.AResponse, err error) {
	article := &model.Article{}
	err = l.svcCtx.DB.QueryRow(article, "select * from article where  uuid = ?", req.Uuid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp = &types.AResponse{
		Code: 200,
		Msg:  "文章具体内容",
		Date: types.Articles{
			Uuid:    article.Uuid,
			UserId:  article.UserId,
			Name:    article.Name,
			Info:    article.Info.String,
			Tag:     article.Tag.String,
			Time:    article.Time.Unix(),
			LikeSum: article.LikeSum.Int64,
		},
	}

	return
}
