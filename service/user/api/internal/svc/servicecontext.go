package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"user/api/internal/config"
	"user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	Redis   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis:   redis.MustNewRedis(c.CacheRedis[0].RedisConf),
	}
}
