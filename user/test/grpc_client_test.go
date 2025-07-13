package test

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"Debate-System/user/client"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	c, err := client.NewUserServiceClientFromConfig("J:\\Debate-System\\user\\client\\user-client.yaml")
	if err != nil {
		panic(err)
	}

	// 4. 调用方法
	resp, err := c.UserLogin(context.Background(), &user_grpc.LoginReq{
		Account:  "test",
		Password: "test",
	})
	assert.NoError(t, err)
	t.Log(resp)
}
