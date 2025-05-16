package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestLog(t *testing.T) {
	logx.SetLevel(logx.ErrorLevel)
	ctx := context.Background()
	logx.WithContext(ctx).Info("hello world")
	logx.Error("test")
	logx.Info(logx.Field("key", "value"))
}
