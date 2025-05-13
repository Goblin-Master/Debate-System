package config

import (
	"Debate-System/pkg/gormx"
	snowflake "Debate-System/utils/snowfake"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	App rest.RestConf
	DB  gormx.Mysql
}

var Node, _ = snowflake.NewNode(1)
