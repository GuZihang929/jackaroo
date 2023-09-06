package logic

import (
	"context"
	"strconv"
	"time"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAccessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAccessLogic {
	return &UserAccessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAccessLogic) UserAccess(req *types.AccessAddRequest) (resp *types.AccessAddResponse, err error) {
	// todo: add your logic here and delete this line

	get, err := l.svcCtx.Redis.Get("user:" + req.Account)
	if len(get) == 0 {
		_, err = l.svcCtx.Redis.Incr("User_Sum:" + strconv.Itoa(time.Now().Hour()))
		if err != nil {
			resp = &types.AccessAddResponse{
				Code: 202,
				Msg:  "人数添加失败",
				Data: "",
			}
		}
		timeout, _ := context.WithTimeout(context.Background(), 5*time.Minute)
		err = l.svcCtx.Redis.SetCtx(timeout, "user:"+req.Account, "1")
		if err != nil {
			resp = &types.AccessAddResponse{
				Code: 202,
				Msg:  "定时失败",
				Data: "",
			}
		}
	}

	resp = &types.AccessAddResponse{
		Code: 200,
		Msg:  "添加成功",
		Data: "",
	}
	return
}
