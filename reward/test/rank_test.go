package test

import (
	"Debate-System/reward/internal/config"
	"Debate-System/reward/internal/repo/cache"
	"Debate-System/reward/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

var configFile = flag.String("f", "..\\etc\\reward-api.yaml", "the config file")

func TestRank(t *testing.T) {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	rdb := cache.NewRewardCache(context.Background(), ctx)
	data := rdb.GetTopN(3)
	logx.Info(data)
}
