package main

import (
	"flag"
	"fmt"

	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/internal/config"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/internal/server"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
