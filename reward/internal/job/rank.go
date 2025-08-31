package job

import (
	"Debate-System/reward/internal/repo/cache"
	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

var (
	CALCULATE_ERROR = errors.New("定时更新排行榜失败")
)

type CalculateRankJob struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	rankCache *cache.RewardCache
	lock      sync.Mutex // 防止替换本地缓存的同时，有人在读数据，发生并发冲突
	n         int
}

func NewCalculateRankJob(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateRankJob {
	return &CalculateRankJob{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		rankCache: cache.NewRewardCache(ctx, svcCtx),
		n:         10, //内置暂时写死
	}
}
func (c *CalculateRankJob) Exec() {
	c.Logger.Infof("开始计算排行榜前%d名", c.n)
	list, err := c.rankCache.CalculateTopN(c.n)
	if err != nil {
		c.Logger.Error(CALCULATE_ERROR)
		return
	}
	data := make([]types.BaseTopNScore, len(list))
	//掉用用户服务去获取基本信息，通过grpc
	for i, v := range list {
		data[i] = types.BaseTopNScore{
			UserID:   v.UserID,
			Score:    v.Score,
			Avatar:   "test",
			Nickname: "test",
		}
	}
	//加锁，防止并发冲突
	c.lock.Lock()
	c.svcCtx.TopN = data
	c.lock.Unlock()
	c.Logger.Infof("计算排行榜任务结束")
}
