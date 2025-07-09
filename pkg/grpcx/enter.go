package grpcx

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type Server struct {
	*grpc.Server
	Port   int
	Name   string
	cancel func()
}

// Serve 启动服务器并且阻塞
func (s *Server) Serve() error {
	// 初始化一个控制整个过程的 ctx
	// 你也可以考虑让外面传进来，这样的话就是 main 函数自己去控制了
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	port := strconv.Itoa(s.Port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// 要先确保启动成功，再注册服务
	err = s.register(ctx, port)
	if err != nil {
		return err
	}
	fmt.Printf("starting grpc server at port %d,name %s\n", s.Port, s.Name)
	return s.Server.Serve(l)
}

func (s *Server) register(ctx context.Context, port string) error {
	return nil
}

func (s *Server) Close() error {
	s.cancel()
	s.Server.GracefulStop()
	return nil
}
