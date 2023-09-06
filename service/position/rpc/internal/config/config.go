package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	UserRpc    zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	CompanyRpc zrpc.RpcClientConf
}
