package config

import (
	"Debate-System/pkg/gormx"
	"Debate-System/pkg/ossx"
	"Debate-System/utils/jwtx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB   gormx.Mysql
	Auth jwtx.Auth
	OSS  ossx.ALiYun
}
