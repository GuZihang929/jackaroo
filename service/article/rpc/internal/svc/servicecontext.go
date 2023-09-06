package svc

import (
	"article/model"
	"article/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	ArtModel model.ArticleModel
	Redis    *redis.Redis
	DB       sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		ArtModel: model.NewArticleModel(conn, c.CacheRedis),
		Redis:    redis.MustNewRedis(c.CacheRedis[0].RedisConf),
		DB:       conn,
	}
}
