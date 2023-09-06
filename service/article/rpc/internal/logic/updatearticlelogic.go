package logic

import (
	"article/rpc/internal/svc"
	"article/rpc/rpc"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"jackaroo/common"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *rpc.UpdateArtRequest) (*rpc.UpdateArtResponse, error) {

	_, err := l.svcCtx.DB.Exec("update article set name = ?,info = ? ,tag =?, time =? where uuid = ?",
		in.Name, in.Info, in.Tag, common.GetTime(in.Time), in.ArtId)

	if err != nil {
		fmt.Println("--------------修改失败", err)
		return nil, err
	}

	return &rpc.UpdateArtResponse{
		Code: 200,
		Msq:  "修改成功",
		Date: nil,
	}, nil
}
