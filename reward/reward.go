package main

import (
	"Debate-System/reward/internal/config"
	"Debate-System/reward/internal/handler"
	"Debate-System/reward/internal/job"
	"Debate-System/reward/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/reward-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//开启定时任务计算排行榜
	job.Cron(context.Background(), ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.HttpServer.Host, c.HttpServer.Port)
	server.Start()
}
