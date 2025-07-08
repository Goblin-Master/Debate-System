package test

import (
	"Debate-System/user/internal/config"
	"Debate-System/user/internal/logic/user"
	"Debate-System/user/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\user\\etc\\user-api.yaml", "the config file")

func TestERROR(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	server := user.NewUserRegisterLogic(context.Background(), ctx)
	server.Logger.Info("123")
}
