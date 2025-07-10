package client

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

// 定义客户端配置结构体
type ClientConfig struct {
	zrpc.RpcClientConf
}

// 初始化 UserServiceClient（支持链路追踪）
func NewUserServiceClientFromConfig(configFile string) (user_grpc.UserServiceClient, error) {
	var c ClientConfig
	conf.MustLoad(configFile, &c)
	zrpcClient := zrpc.MustNewClient(c.RpcClientConf)
	return user_grpc.NewUserServiceClient(zrpcClient.Conn()), nil
}

// 使用示例
func Example() {
	client, err := NewUserServiceClientFromConfig("J:\\Debate-System\\user\\client\\user-client.yaml")
	if err != nil {
		panic(err)
	}
	// 现在 client 就可以正常调用 gRPC 方法，并自动传递 trace/span
	_ = client
}
