package config

import (
	snowflake "Debate-System/pkg/snowfake"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
}

var Node, _ = snowflake.NewNode(1)
