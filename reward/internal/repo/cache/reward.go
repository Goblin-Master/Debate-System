package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RewardCache struct {
	ctx context.Context
	rdb redis.Cmdable
}

func NewRewardCache(ctx context.Context, rdb redis.Cmdable) *RewardCache {
	return &RewardCache{
		ctx: ctx,
		rdb: rdb,
	}
}
