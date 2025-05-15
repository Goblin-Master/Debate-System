package config

import (
	"Debate-System/pkg/gormx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	App rest.RestConf
	DB  gormx.Mysql
}
