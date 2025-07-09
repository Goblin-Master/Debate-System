package test

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"Debate-System/user/client"
	"context"
	"github.com/stretchr/testify/assert"
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
