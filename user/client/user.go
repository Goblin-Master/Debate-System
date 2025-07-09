package client

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"context"
	"google.golang.org/grpc"
)

type UserClient struct {
	client user_grpc.UserServiceClient
}

func NewUserClient(c user_grpc.UserServiceClient) *UserClient {
	return &UserClient{
		client: c,
	}
}
func (u *UserClient) UserLogin(ctx context.Context, in *user_grpc.LoginReq, opts ...grpc.CallOption) (*user_grpc.LoginResp, error) {
	return u.client.UserLogin(ctx, in, opts...)
}

func (u *UserClient) UserRegister(ctx context.Context, in *user_grpc.UserRegisterReq, opts ...grpc.CallOption) (*user_grpc.UserRegisterResp, error) {
	return u.client.UserRegister(ctx, in, opts...)
}

func (u *UserClient) UserInfo(ctx context.Context, in *user_grpc.UserInfoReq, opts ...grpc.CallOption) (*user_grpc.UserInfoResp, error) {
	return u.client.UserInfo(ctx, in, opts...)
}

func (u *UserClient) UserModify(ctx context.Context, in *user_grpc.UserModifyReq, opts ...grpc.CallOption) (*user_grpc.UserModifyResp, error) {
	return u.client.UserModify(ctx, in, opts...)
}
