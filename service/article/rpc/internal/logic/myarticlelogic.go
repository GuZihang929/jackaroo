package logic

import (
	"article/model"
	"article/rpc/internal/svc"
	"article/rpc/rpc"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyArticleLogic {
	return &MyArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MyArticleLogic) MyArticle(in *rpc.MyArtRequest) (*rpc.MyArtResponse, error) {
	var Arts []model.Article

	err := l.svcCtx.DB.QueryRows(&Arts, "select * from article where user_id = ?", in.ArtId)
	if err != nil {
		fmt.Println("--------------", "查询我的文章出错")
		return nil, err
	}
	resp := &rpc.MyArtResponse{
		Code: 200,
		Date: []*rpc.Art{},
	}

	for _, date := range Arts {
		resp.Date = append(resp.Date, &rpc.Art{
			Uuid:    date.Uuid,
			UserId:  date.UserId,
			Name:    date.Name,
			Info:    date.Info.String,
			Tag:     date.Tag.String,
			Time:    date.Time.Unix(),
			LikeSum: date.LikeSum.Int64,
		})
	}

	return resp, err
}
