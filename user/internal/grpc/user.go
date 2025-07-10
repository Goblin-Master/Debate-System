package grpc

import (
	user_grpc "Debate-System/api/proto/gen/user"
	"Debate-System/user/internal/logic/user"
	"Debate-System/user/internal/svc"
	"Debate-System/user/internal/types"
	"context"
	"google.golang.org/grpc"
)

type UserServiceServer struct {
	// 正常我都会组合这个
	user_grpc.UnimplementedUserServiceServer
	svcCtx *svc.ServiceContext
}

func NewUserServiceServer(serverCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: serverCtx,
	}
}

func (u *UserServiceServer) Register(server grpc.ServiceRegistrar) {
	user_grpc.RegisterUserServiceServer(server, u)
}

func (u *UserServiceServer) UserLogin(ctx context.Context, req *user_grpc.LoginReq) (*user_grpc.LoginResp, error) {
	l := user.NewUserLoginLogic(ctx, u.svcCtx)
	resp, err := l.UserLogin(&types.LoginReq{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &user_grpc.LoginResp{
		Nickname: resp.Nickname,
		UserId:   resp.UserID,
		Avatar:   resp.Avatar,
		Token:    resp.Token,
	}, nil
}

func (u *UserServiceServer) UserRegister(ctx context.Context, req *user_grpc.UserRegisterReq) (*user_grpc.UserRegisterResp, error) {
	l := user.NewUserRegisterLogic(ctx, u.svcCtx)
	resp, err := l.UserRegister(&types.UserRegisterReq{
		Account:  req.Account,
		Password: req.Password,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	return &user_grpc.UserRegisterResp{
		UserId: resp.UserID,
	}, nil
}

func (u *UserServiceServer) UserInfo(ctx context.Context, req *user_grpc.UserInfoReq) (*user_grpc.UserInfoResp, error) {
	l := user.NewUserInfoLogic(ctx, u.svcCtx)
	resp, err := l.UserInfo(&types.UserInfoReq{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &user_grpc.UserInfoResp{
		Nickname: resp.Nickname,
		Avatar:   resp.Avatar,
	}, nil
}

func (u *UserServiceServer) UserModify(ctx context.Context, req *user_grpc.UserModifyReq) (*user_grpc.UserModifyResp, error) {
	l := user.NewUserModifyLogic(ctx, u.svcCtx)
	resp, err := l.UserModify(&types.UserModifyReq{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		UserID:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &user_grpc.UserModifyResp{
		Message: resp.Message,
	}, nil
}
