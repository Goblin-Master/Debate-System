package score

import (
	"context"

	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopNScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopNScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopNScoreLogic {
	return &TopNScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopNScoreLogic) TopNScore(req *types.TopNScoreReq) (resp *types.TopNScoreResp, err error) {
	// todo: add your logic here and delete this line

	return
}
