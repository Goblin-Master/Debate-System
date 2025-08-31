package cache

import (
	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"
	"context"

	"strconv"
)

const RANK = "rank"

type RewardCache struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRewardCache(ctx context.Context, svcCtx *svc.ServiceContext) *RewardCache {
	return &RewardCache{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type TopN struct {
	UserID int64
	Score  int
}

func (rc *RewardCache) CalculateTopN(n int) ([]TopN, error) {
	// 从 Redis 获取前 N 个元素（按分数从高到低）
	zs, err := rc.svcCtx.RDB.ZRevRangeWithScores(rc.ctx, RANK, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}

	// 转成你自己的结构体
	topN := make([]TopN, 0, len(zs))
	for _, z := range zs {
		userID, err := strconv.ParseInt(z.Member.(string), 10, 64)
		if err != nil {
			continue // 或返回错误
		}
		topN = append(topN, TopN{
			UserID: userID,
			Score:  int(z.Score),
		})
	}
	return topN, nil
}

func (rc *RewardCache) ModifyScore(user_id int64, score int) error {
	// 1. 把本次增量写进 Redis 的 zset
	//    如果 member 已存在，ZINCRBY 会在原 score 上累加
	if err := rc.svcCtx.RDB.ZIncrBy(
		rc.ctx,
		RANK,           // zset 的 key
		float64(score), // 要增加的分数
		strconv.FormatInt(user_id, 10),
	).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *RewardCache) GetTopN(n int) []types.BaseTopNScore {
	length := len(rc.svcCtx.TopN)
	if length == 0 {
		return []types.BaseTopNScore{}
	}
	cnt := max(length, n)
	result := rc.svcCtx.TopN[:cnt]
	return result
}
