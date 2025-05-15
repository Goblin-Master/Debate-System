package user

import (
	"Debate-System/internal/repo"
	"context"

	"Debate-System/internal/svc"
	"Debate-System/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   *repo.UserRepo
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repo.NewUserRepo(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	id, err := l.repo.CreateUser(req.Account, req.Password, req.Nickname)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return &types.UserRegisterResp{
		UserID: id,
	}, nil
}
