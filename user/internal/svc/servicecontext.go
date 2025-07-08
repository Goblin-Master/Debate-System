package svc

import (
	"Debate-System/user/internal/config"
	"Debate-System/user/internal/model"
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
		DB:     model.InitDB(c.DB, nil),
		Node:   snowflake.SetNode(1),
	}
}
