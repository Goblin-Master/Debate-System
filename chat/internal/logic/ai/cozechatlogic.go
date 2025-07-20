package ai

import (
	"context"

	"Debate-System/chat/internal/svc"
	"Debate-System/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CozeChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCozeChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CozeChatLogic {
	return &CozeChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CozeChatLogic) CozeChat(req *types.CozeChatReq) (resp *types.CozeChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
