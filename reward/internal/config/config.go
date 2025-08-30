package config

import (
	"Debate-System/pkg/gormx"
	"Debate-System/pkg/redisx"
	"Debate-System/utils/jwtx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// 将 HTTP 服务配置放入一个独立的字段，避免嵌入冲突
	HttpServer rest.RestConf

	// 为 gRPC 服务配置一个独立的字段
	RpcServer zrpc.RpcServerConf

	// 在顶层显式定义共享的 Auth 配置，解决冲突问题
	Auth jwtx.Auth

	// 3. 公共配置
	DB  gormx.Mysql
	RDB redisx.Redis
}
