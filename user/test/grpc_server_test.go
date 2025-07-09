package test

import (
	"Debate-System/pkg/grpcx"
	"Debate-System/user/internal/config"
	mygrpc "Debate-System/user/internal/grpc"
	"Debate-System/user/internal/svc"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\user\\etc\\user-api.yaml", "the config file")

func TestServer(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	ctx := svc.NewServiceContext(c)
	user := mygrpc.NewUserServiceServer(ctx)
	server := grpc.NewServer()
	user.Register(server)
	s := &grpcx.Server{
		Name:   "user-service",
		Port:   8080,
		Server: server,
	}
	err := s.Serve()
	if err != nil {
		panic(err)
	}
}
