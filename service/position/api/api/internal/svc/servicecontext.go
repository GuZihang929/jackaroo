package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jackaroo/service/position/api/api/internal/config"
	"jackaroo/service/position/rpc/positionclient"
	"jackaroo/service/position/rpc/types/position"
)

type ServiceContext struct {
	Config      config.Config
	PositionRpc position.PositionClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		PositionRpc: positionclient.NewPosition(zrpc.MustNewClient(c.PositionRpc)),
	}
}
