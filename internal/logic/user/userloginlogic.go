package user

import (
	"Debate-System/internal/repo"
	"Debate-System/utils/jwtx"
	"context"
	"errors"

	"Debate-System/internal/svc"
	"Debate-System/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   *repo.UserRepo
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repo.NewUserRepo(ctx, svcCtx),
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.repo.CheckLogin(req.Account, req.Password)
	if err != nil {
		if !errors.Is(err, repo.ACCOUNT_OR_PWD_ERROR) {
			l.Logger.Error(err)
		}
		return nil, ACCOUNT_OR_PWD_ERROR
	}

	token, err := jwtx.GenToken(jwtx.Claims{
		UserID: user.UserID,
		Auth: jwtx.Auth{
			AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
			AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		},
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, DEFAULT_ERROR
	}

	resp = &types.LoginResp{
		Nickname: user.Nickname,
		UserID:   user.UserID,
		Avatar:   user.Avatar,
		Token:    token,
	}
	return resp, nil
}
