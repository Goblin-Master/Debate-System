package config

import (
	"Debate-System/pkg/gormx"
	"Debate-System/utils/jwtx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	App  rest.RestConf
	DB   gormx.Mysql
	Auth jwtx.Auth
}
