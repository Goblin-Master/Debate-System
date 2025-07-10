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
	resp, err := c.UserInfo(context.Background(), &user_grpc.UserInfoReq{
		UserId: 1941033845134462976,
	})
	assert.NoError(t, err)
	t.Log(resp)
}
