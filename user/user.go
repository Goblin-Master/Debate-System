package main

import (
	"Debate-System/user/internal/config"
	mygrpc "Debate-System/user/internal/grpc"
	"Debate-System/user/internal/handler"
	"Debate-System/user/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	ctx := svc.NewServiceContext(c)

	go func() {
		user := mygrpc.NewUserServiceServer(ctx)

		server := zrpc.MustNewServer(c.RpcServer, func(grpcServer *grpc.Server) {
			user.Register(grpcServer)
		})
		defer server.Stop()
		fmt.Printf("Starting rpc server at %s...\n", c.RpcServer.ListenOn)
		server.Start()
	}()

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting http server at %s:%d...\n", c.HttpServer.Host, c.HttpServer.Port)
	server.Start()
}
