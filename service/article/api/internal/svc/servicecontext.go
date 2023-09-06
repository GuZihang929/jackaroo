package svc

import (
	"article/api/internal/config"
	"article/rpc/article"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.Article
	DB         sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.ArticleRpc)
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
		DB:         sqlx.NewMysql(c.Mysql.DataSource),
	}
}
