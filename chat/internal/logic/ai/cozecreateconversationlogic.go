package ai

import (
	"context"

	"Debate-System/chat/internal/svc"
	"Debate-System/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CozeCreateConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCozeCreateConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CozeCreateConversationLogic {
	return &CozeCreateConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CozeCreateConversationLogic) CozeCreateConversation(req *types.CozeCreateConversationReq) (resp *types.CozeCreateConversationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
