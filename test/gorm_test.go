package test

import (
	"Debate-System/internal/config"
	"Debate-System/internal/logic/user"
	"Debate-System/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\etc\\app-api.yaml", "the config file")

func TestERROR(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	server := user.NewUserRegisterLogic(context.Background(), ctx)
	server.Logger.Info("123")
}
