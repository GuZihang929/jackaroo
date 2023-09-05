package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"xiangxiang/jackaroo/service/company/rpc/companyclient"
	"xiangxiang/jackaroo/service/company/rpc/types/company"
	"xiangxiang/jackaroo/service/position/model/model"
	"xiangxiang/jackaroo/service/position/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	PositionModel model.PositionModel
	CompanyRpc    company.CompanyClient
	DB            sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		PositionModel: model.NewPositionModel(conn, c.CacheRedis),
		CompanyRpc:    companyclient.NewCompany(zrpc.MustNewClient(c.CompanyRpc)),
		DB:            conn,
	}
}
