package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"job/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	DB     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.CacheRedis[0].RedisConf),
		DB:     conn,
	}
}
