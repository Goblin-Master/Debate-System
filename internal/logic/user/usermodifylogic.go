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

type UserModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   *repo.UserRepo
}

func NewUserModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModifyLogic {
	return &UserModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repo.NewUserRepo(ctx, svcCtx),
	}
}

func (l *UserModifyLogic) UserModify(req *types.UserModifyReq) (resp *types.UserModifyResp, err error) {
	if req.Avatar == "" && req.Nickname == "" {
		return nil, VAILD_EMPTY
	}
	user_id, err := jwtx.GetUserID(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return nil, DEFAULT_ERROR
	}
	err = l.repo.ModifyUserData(user_id, req.Nickname, req.Avatar)
	if err != nil {
		if errors.Is(err, repo.USER_NOT_EXIST) {
			return nil, USER_NOT_EXIST
		}
		l.Logger.Error(err)
		return nil, MODIFY_ERROR
	}
	resp = &types.UserModifyResp{Message: "修改成功"}
	return resp, nil
}
