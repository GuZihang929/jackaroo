package logic

import (
	"article/model"
	"article/rpc/internal/svc"
	"article/rpc/rpc"
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"jackaroo/common"
)

type AddMyArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMyArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMyArticleLogic {
	return &AddMyArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMyArticleLogic) AddMyArticle(in *rpc.AddArtRequest) (*rpc.AddArtResponse, error) {

	_, err := l.svcCtx.ArtModel.Insert(l.ctx, &model.Article{
		Uuid:    common.NewSnowflake().NextVal(),
		UserId:  in.UserId,
		Name:    in.Name,
		Info:    sql.NullString{String: in.Info, Valid: true},
		Tag:     sql.NullString{String: in.Tag, Valid: true},
		Time:    common.GetTime(in.Time),
		LikeSum: sql.NullInt64{Int64: 0},
	})
	if err != nil {
		return &rpc.AddArtResponse{
			Code: 200,
			Msq:  "添加失败",
		}, err
	}

	return &rpc.AddArtResponse{
		Code: 200,
		Msq:  "添加成功",
	}, err
}
