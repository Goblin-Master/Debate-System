package user

import (
	"Debate-System/internal/repo"
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
		repo:   repo.NewUserRepo(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.repo.CheckLogin(req.Account, req.Password)
	if err != nil {
		if !errors.Is(err, repo.ACCOUNT_OR_PWD_ERROR) {
			l.Logger.Error(err)
		}
		return nil, err
	}
	resp = &types.LoginResp{
		Nickname: user.Nickname,
		UserID:   user.UserID,
		Avatar:   user.Avatar,
	}
	return resp, nil
}
