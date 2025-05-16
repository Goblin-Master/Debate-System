package svc

import (
	"Debate-System/internal/config"
	"Debate-System/pkg/gormx"
	snowflake "Debate-System/utils/snowfake"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Node   *snowflake.Node
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     gormx.InitDB(c.DB, nil),
		Node:   snowflake.SetNode(1),
	}
}
