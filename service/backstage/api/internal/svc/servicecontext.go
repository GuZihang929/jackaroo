package svc

import (
	"backstage/api/internal/config"
	"backstage/rpc/backstage"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	BackstageRpc backstage.Backstage
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		BackstageRpc: backstage.NewBackstage(zrpc.MustNewClient(c.BackstageRpc)),
	}
}
