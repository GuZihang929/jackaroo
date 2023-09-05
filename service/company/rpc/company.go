package main

import (
	"flag"
	"fmt"

	"xiangxiang/jackaroo/service/company/rpc/internal/config"
	"xiangxiang/jackaroo/service/company/rpc/internal/server"
	"xiangxiang/jackaroo/service/company/rpc/internal/svc"
	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/company.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		company.RegisterCompanyServer(grpcServer, server.NewCompanyServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
