package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jackaroo/service/company/model/model"
	"jackaroo/service/company/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	CompanyModel model.CompanyModel
	DB           sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		CompanyModel: model.NewCompanyModel(conn, c.CacheRedis),
		DB:           conn,
	}
}
