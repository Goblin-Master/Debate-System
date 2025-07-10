package test

import (
	"Debate-System/user/internal/config"
	mygrpc "Debate-System/user/internal/grpc"
	"Debate-System/user/internal/svc"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"

	"google.golang.org/grpc"
	"testing"
)

// 无需在意错误，换掉Config就不会报错了
var configFile = flag.String("f", "J:\\Debate-System\\user\\etc\\user-rpc.yaml", "the config file")

func TestServer(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	ctx := svc.NewServiceContext(c)
	user := mygrpc.NewUserServiceServer(ctx)

	// 用 go-zero zrpc 启动 gRPC 服务，trace 自动集成
	server := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.Register(grpcServer)
	})
	defer server.Stop()
	server.Start()
}
