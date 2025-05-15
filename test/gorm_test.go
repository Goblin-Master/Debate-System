package test

import (
	"Debate-System/internal/config"
	"Debate-System/internal/global"
	"Debate-System/internal/repo/dao"
	"Debate-System/pkg/gormx"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/gorm"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\etc\\app-api.yaml", "the config file")

func TestERROR(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	global.DB = gormx.MustOpen(c.DB, nil)
	t.Log(global.DB)
	v := dao.NewUserDao()
	t.Log(v)
	_, err := v.GetByID(1)
	t.Log(err == gorm.ErrRecordNotFound)
}
