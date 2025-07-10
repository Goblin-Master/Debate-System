package config

import (
	"Debate-System/pkg/gormx"
	"Debate-System/pkg/ossx"
	"Debate-System/utils/jwtx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB   gormx.Mysql
	Auth jwtx.Auth
	OSS  ossx.ALiYun
}

// 这个用于rpc服务的config，改Auth是因为冲突了
type RpcConfig struct {
	zrpc.RpcServerConf
	DB      gormx.Mysql
	RpcAuth jwtx.Auth
	OSS     ossx.ALiYun
}
