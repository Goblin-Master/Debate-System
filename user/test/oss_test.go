package test

import (
	"Debate-System/pkg/ossx"
	"context"
	"testing"
)

func TestOSS(t *testing.T) {
	ossx.InitALiYunOSS(context.Background(), ossx.ALiYun{
		Enable:    true,
		AccessID:  "",
		AccessKey: "",
		Region:    "oss-cn-beijing",
	})
}
