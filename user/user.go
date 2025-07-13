package main

import (
	"Debate-System/user/internal/config"
	"Debate-System/user/internal/handler"
	"Debate-System/user/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.HttpServer.Host, c.HttpServer.Port)
	server.Start()
}
