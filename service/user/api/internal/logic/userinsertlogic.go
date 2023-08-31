package logic

import (
	"context"
	"encoding/json"
	"user/rpc/rpc"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInsertLogic {
	return &UserInsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInsertLogic) UserInsert(req *types.UserInsertRequst) (*types.UserInsertRespinse, error) {

	i, err := l.ctx.Value("userId").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInsert(l.ctx, &rpc.UserInsertRequest{
		Name:     req.Name,
		Age:      int64(req.Age),
		Sex:      req.Sex,
		Identity: req.Identity,
		Need:     req.Need,
		Uuid:     i,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserInsertRespinse{
		Code: int(res.Code),
		Msg:  res.Msg,
	}, nil
}
