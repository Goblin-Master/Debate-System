package score

import (
	"Debate-System/reward/internal/repo"
	"context"

	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopNScoreLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	rewardRepo *repo.RewardRepo
}

func NewTopNScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopNScoreLogic {
	return &TopNScoreLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		rewardRepo: repo.NewRewardRepo(ctx, svcCtx),
	}
}

func (l *TopNScoreLogic) TopNScore(req *types.TopNScoreReq) (resp *types.TopNScoreResp, err error) {
	resp = &types.TopNScoreResp{
		List: l.rewardRepo.GetTopN(req.N),
	}
	return resp, nil
}
