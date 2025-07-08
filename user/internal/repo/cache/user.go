package cache

import "Debate-System/user/internal/svc"

type UserCache struct {
	svcCtx *svc.ServiceContext
}

func NewUserCache(svcCtx *svc.ServiceContext) *UserCache {
	return &UserCache{
		svcCtx: svcCtx,
	}
}
