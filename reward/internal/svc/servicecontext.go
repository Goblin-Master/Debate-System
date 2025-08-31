package svc

import (
	"Debate-System/pkg/redisx"
	"Debate-System/reward/internal/config"
	"Debate-System/reward/internal/model"
	"Debate-System/reward/internal/types"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	RDB    redis.Cmdable
	DB     *gorm.DB
	TopN   []types.BaseTopNScore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RDB:    redisx.InitRedis(c.RDB),
		DB:     model.InitDB(c.DB, nil),
		TopN:   make([]types.BaseTopNScore, 0),
	}
}
