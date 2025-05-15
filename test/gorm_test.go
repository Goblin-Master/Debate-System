package test

import (
	"Debate-System/internal/config"
	"Debate-System/internal/global"
	"Debate-System/internal/logic/user"
	"Debate-System/internal/svc"
	"Debate-System/pkg/gormx"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\etc\\app-api.yaml", "the config file")

func TestERROR(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	global.DB = gormx.MustOpen(c.DB, nil)
	ctx := svc.NewServiceContext(c)
	server := user.NewUserRegisterLogic(context.Background(), ctx)
	server.UserRegister(nil)
}
