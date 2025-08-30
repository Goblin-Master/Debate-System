package repo

import (
	"Debate-System/reward/internal/repo/cache"
	"Debate-System/reward/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type IRewardRepo interface {
}
type RewardRepo struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	rewardCache *cache.RewardCache
}

var _ IRewardRepo = (*RewardRepo)(nil)

func NewRewardRepo(ctx context.Context, svcCtx *svc.ServiceContext) *RewardRepo {
	return &RewardRepo{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		rewardCache: cache.NewRewardCache(ctx, svcCtx.RDB),
	}
}
