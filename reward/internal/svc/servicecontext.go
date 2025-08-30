package svc

import (
	"Debate-System/pkg/redisx"
	"Debate-System/reward/internal/config"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config config.Config
	RDB    redis.Cmdable
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RDB:    redisx.InitRedis(c.RDB),
	}
}
