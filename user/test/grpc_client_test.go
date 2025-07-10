package test

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"Debate-System/user/client"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	// 1. 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	assert.NoError(t, err)
	defer conn.Close()
	// 2. 创建 gRPC 客户端
	grpcClient := user_grpc.NewUserServiceClient(conn)
	// 3. 包装为你自己的 UserClient
	userClient := client.NewUserClient(grpcClient)

	// 4. 调用方法
	resp, err := userClient.UserLogin(context.Background(), &user_grpc.LoginReq{
		Account:  "test",
		Password: "test",
	})
	assert.NoError(t, err)
	t.Log(resp)
}
func initClient(ctx context.Context) *client.UserClient {
	// --------------------【核心改造区域开始】--------------------

	// 1. 定义 zrpc 客户端配置
	var clientConf zrpc.RpcClientConf

	// 2. 为配置填充默认值
	// 这一步会自动启用 OpenTelemetry, 确保 trace.UnaryClientInterceptor() 在内部被使用
	conf.FillDefault(&clientConf)

	// 3. 设置目标服务的地址
	// 对于直连，推荐使用 "direct:///ip:port" 格式
	clientConf.Target = "direct:///localhost:8080"
	// 如果有 etcd，则使用 clientConf.Etcd.Hosts 和 clientConf.Etcd.Key

	// 4. 使用 go-zero 的方式创建客户端，它会自动处理连接和所有拦截器
	zrpcClient := zrpc.MustNewClient(clientConf)

	// --------------------【核心改造区域结束】--------------------

	// ---- 后续业务逻辑几乎不变 ----

	// 从 zrpc.Client 中获取原始的 gRPC 连接，并创建业务客户端
	grpcClient := user_grpc.NewUserServiceClient(zrpcClient.Conn())
	userClient := client.NewUserClient(grpcClient)
	return userClient
}
