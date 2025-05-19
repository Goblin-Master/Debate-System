package user

import (
	"context"

	"Debate-System/internal/svc"
	"Debate-System/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAutoLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAutoLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAutoLoginLogic {
	return &UserAutoLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAutoLoginLogic) UserAutoLogin() (resp *types.UserAutoLoginResp, err error) {
	resp = &types.UserAutoLoginResp{
		Message: "用户自动登录成功",
	}
	return resp, nil
}
