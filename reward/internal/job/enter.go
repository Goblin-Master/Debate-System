package job

import (
	"Debate-System/reward/internal/svc"
	"context"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func Cron(ctx context.Context, svcCtx *svc.ServiceContext) {
	l := logx.WithContext(ctx)
	zone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		l.Error(err)
	}
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(zone))
	rankJob := NewCalculateRankJob(ctx, svcCtx)
	// 每隔2分钟计算一下排行榜
	_, err = crontab.AddFunc("0 */2 * * * *", rankJob.Exec)
	if err != nil {
		l.Error(err)
	}
	crontab.Start()
}
