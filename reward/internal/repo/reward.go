package repo

import (
	"Debate-System/reward/internal/repo/cache"
	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	DEFAULT_ERROR = errors.New("默认错误")
)

type IRewardRepo interface {
	GetTopN(n int) []types.BaseTopNScore
	ModifyScore(user_id int64, score int) error
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
		rewardCache: cache.NewRewardCache(ctx, svcCtx),
	}
}

func (r *RewardRepo) GetTopN(n int) []types.BaseTopNScore {
	result := r.rewardCache.GetTopN(n)
	return result
}

func (r *RewardRepo) ModifyScore(user_id int64, score int) error {
	err := r.rewardCache.ModifyScore(user_id, score)
	if err != nil {
		r.Logger.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}
