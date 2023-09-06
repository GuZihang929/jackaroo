package logic

import (
	"context"
	"jackaroo/service/company/model/model"

	"jackaroo/service/company/rpc/internal/svc"
	"jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *company.RefreshRequest) (*company.RefreshResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	db, err := l.svcCtx.DB.RawDB()
	if err != nil {
		return nil, err
	}
	// 准备查询语句
	query := "SELECT COUNT(*) FROM position WHERE company = ?"
	// 执行查询并获取结果
	var count int64
	err = db.QueryRow(query, res.Company).Scan(&count)
	if err != nil {
		return nil, err
	}

	update := "UPDATE company SET number = ? WHERE id =?"
	var com model.Company
	err = db.QueryRow(update, count, res.Id).Err()
	if err != nil {
		return nil, err
	}
	return &company.RefreshResponse{
		Id:      com.Id,
		Company: res.Company,
		Number:  count,
		Brief:   res.Brief,
		Picture: res.Picture,
	}, nil
}
