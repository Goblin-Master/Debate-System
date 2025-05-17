package ossx

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
)

type QiniuService struct {
	logx.Logger
	client     *uploader.UploadManager
	ctx        context.Context
	bucketName string
	Prefix     string
	url        string
}

var _ Service = (*QiniuService)(nil)

func NewQiNiuOSSService(ctx context.Context, client *uploader.UploadManager, bucketName, prefix, url string) *QiniuService {
	return &QiniuService{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		client:     client,
		bucketName: bucketName,
		Prefix:     prefix,
		url:        url,
	}
}

func (s *QiniuService) UploadFile(file *multipart.FileHeader) (string, error) {
	key := fmt.Sprintf("%s/%s", s.Prefix, file.Filename)
	reader, err := file.Open()
	if err != nil {
		s.Logger.Error(err)
		return "", FILE_READ_FAIL
	}
	defer reader.Close()
	err = s.client.UploadReader(s.ctx, reader, &uploader.ObjectOptions{
		BucketName: s.bucketName,
		ObjectName: &key,
		FileName:   file.Filename,
	}, nil)
	if err != nil {
		s.Logger.Error(err)
		return "", UPLOUD_FILE_FAIL
	}
	return fmt.Sprintf("%s/%s", s.url, key), nil
}
