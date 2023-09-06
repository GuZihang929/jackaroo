package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"job/api/internal/config"
	"job/rpc/sjob"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	SJobRpc sjob.SJob
	DB      sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SJobRpc: sjob.NewSJob(zrpc.MustNewClient(c.SJobRpc)),
		Redis:   redis.MustNewRedis(c.CacheRedis[0].RedisConf),
		DB:      sqlx.NewMysql(c.Mysql.DataSource),
	}
}
