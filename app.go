package main

import (
	"flag"
	"fmt"

	"Debate-System/internal/config"
	"Debate-System/internal/handler"
	"Debate-System/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.App)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Mysql connect successfully %v\n", c.DB)
	fmt.Printf("Starting server at %s:%d...\n", c.App.Host, c.App.Port)

	server.Start()
}
