package logic

import (
	"context"
	"fmt"

	"job/api/internal/svc"
	"job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyLogic {
	return &CompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompanyLogic) Company() (resp *types.CompanySumResponse, err error) {
	fmt.Println("0000000000000")
	var com = []types.Companys{}
	err = l.svcCtx.DB.QueryRows(&com, "select * from company")
	fmt.Println(com)
	if err != nil {
		resp = &types.CompanySumResponse{
			Code: 202,
			Msg:  "失败，重新请求",
			Date: nil,
		}
	}
	resp = &types.CompanySumResponse{
		Code: 200,
		Msg:  "成功",
		Date: com,
	}
	return
}
