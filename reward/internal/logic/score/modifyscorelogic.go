package score

import (
	"context"

	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyScoreLogic {
	return &ModifyScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyScoreLogic) ModifyScore(req *types.ModifyScoreReq) (resp *types.ModifyScoreResp, err error) {
	// todo: add your logic here and delete this line

	return
}
