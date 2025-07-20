package svc

import (
	"Debate-System/chat/internal/config"
	"Debate-System/pkg/cozex"
	"github.com/coze-dev/coze-go"
)

type ServiceContext struct {
	Config     config.Config
	CozeClient coze.CozeAPI
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CozeClient: cozex.InitCozeClient(c.Coze),
	}
}
