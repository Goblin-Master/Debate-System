package oss

import (
	"Debate-System/pkg/ossx"
	"Debate-System/user/internal/repo"
	"Debate-System/user/internal/svc"
	"Debate-System/user/internal/types"
	"context"
	"mime/multipart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   *repo.UserRepo
	oss    ossx.Service
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repo.NewUserRepo(ctx, svcCtx),
		oss:    ossx.InitSSOService(ctx, svcCtx.Config.OSS),
	}
}

func (l *UploadFileLogic) UploadFile(file *multipart.FileHeader) (resp *types.UploadFileResp, err error) {
	url, err := l.oss.UploadFile(file)
	resp = &types.UploadFileResp{
		Url: url,
	}
	return
}
