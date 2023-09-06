package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jackaroo/service/company/api/api/internal/config"
	"jackaroo/service/company/rpc/companyclient"
	"jackaroo/service/company/rpc/types/company"
)

type ServiceContext struct {
	Config     config.Config
	CompanyRpc company.CompanyClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CompanyRpc: companyclient.NewCompany(zrpc.MustNewClient(c.CompanyRpc)),
	}
}
