package logic

import (
	"article/model"
	"context"
	"fmt"

	"article/rpc/internal/svc"
	"article/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticlesLogic) Articles(in *rpc.ArtsRequest) (*rpc.ArtsResponse, error) {

	var Arts []model.Article

	err := l.svcCtx.DB.QueryRows(&Arts, "select * from article where uuid > ? limit ?", int(in.NextID), int(in.PageSize))
	if err != nil {
		fmt.Println("查询文章..............", err)
		return nil, err
	}
	fmt.Println(Arts)
	resp := &rpc.ArtsResponse{
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
