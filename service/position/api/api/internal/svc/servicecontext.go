package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xiangxiang/jackaroo/service/position/api/api/internal/config"
	"xiangxiang/jackaroo/service/position/rpc/positionclient"
	"xiangxiang/jackaroo/service/position/rpc/types/position"
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
