package user

import (
	"Debate-System/user/internal/repo"
	"Debate-System/user/internal/svc"
	"Debate-System/user/internal/types"
	"Debate-System/utils/jwtx"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   *repo.UserRepo
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repo.NewUserRepo(ctx, svcCtx),
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	if req.UserID == 0 {
		req.UserID, err = jwtx.GetUserID(l.ctx)
		if err != nil {
			l.Logger.Error(err)
			return nil, USER_NOT_EXIST
		}
	}
	user, err := l.repo.GetUserByID(req.UserID)
	if err != nil {
		if !errors.Is(err, repo.USER_NOT_EXIST) {
			l.Logger.Error(err)
		}
		return nil, USER_NOT_EXIST
	}
	resp = &types.UserInfoResp{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	return resp, nil
}
