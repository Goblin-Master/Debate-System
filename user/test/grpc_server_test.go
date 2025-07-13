package test

import (
	"Debate-System/user/internal/config"
	mygrpc "Debate-System/user/internal/grpc"
	"Debate-System/user/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"testing"
)

// 无需在意错误，换掉Config就不会报错了
var configFile = flag.String("f", "J:\\Debate-System\\user\\etc\\user-api.yaml", "the config file")

func TestServer(t *testing.T) {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	ctx := svc.NewServiceContext(c)
	user := mygrpc.NewUserServiceServer(ctx)

	server := zrpc.MustNewServer(c.RpcServer, func(grpcServer *grpc.Server) {
		user.Register(grpcServer)
	})
	defer server.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.RpcServer.ListenOn)
	server.Start()
}
