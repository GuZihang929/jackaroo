package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"jackaroo/service/company/rpc/companyclient"
	"jackaroo/service/company/rpc/types/company"
	"jackaroo/service/position/model/model"
	"jackaroo/service/position/rpc/internal/config"
	"user/rpc/user"
)

type ServiceContext struct {
	Config        config.Config
	PositionModel model.PositionModel
	UserRpc       user.User
	CompanyRpc    company.CompanyClient
	DB            sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		PositionModel: model.NewPositionModel(conn, c.CacheRedis),
		UserRpc:       userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CompanyRpc:    companyclient.NewCompany(zrpc.MustNewClient(c.CompanyRpc)),
		DB:            conn,
	}
}
