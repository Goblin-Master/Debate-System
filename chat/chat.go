package main

import (
	"Debate-System/chat/ws"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"Debate-System/chat/internal/config"
	"Debate-System/chat/internal/handler"
	"Debate-System/chat/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	ctx := svc.NewServiceContext(c)

	go ws.Server(c, ctx)

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting http server at %s:%d...\n", c.HttpServer.Host, c.HttpServer.Port)
	server.Start()
}
