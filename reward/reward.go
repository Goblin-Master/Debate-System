package main

import (
	"flag"
	"fmt"

	"Debate-System/reward/internal/config"
	"Debate-System/reward/internal/handler"
	"Debate-System/reward/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/reward-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Println("%#v\n", c)

	fmt.Printf("Starting server at %s:%d...\n", c.HttpServer.Host, c.HttpServer.Port)
	server.Start()
}
