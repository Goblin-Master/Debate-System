package ossx

import (
	"context"
	"errors"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"mime/multipart"
)

var (
	FILE_READ_FAIL   = errors.New("文件读取失败")
	UPLOUD_FILE_FAIL = errors.New("上传文件失败")
)

type QiNiu struct {
	Enable    bool   `json:"enable"`
	AccessKey string `json:"accessKey"`
	Bucket    string `json:"bucket"`
	SecretKey string `json:"secretKey"`
	Url       string `json:"url"`
	Region    string `json:"region"`
	Prefix    string `json:"prefix"`
}

type Service interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func InitSSOService(ctx context.Context, c QiNiu) Service {
	return InitQiNiuOSS(ctx, c)
}

func InitQiNiuOSS(ctx context.Context, c QiNiu) Service {
	mac := credentials.NewCredentials(c.AccessKey, c.SecretKey)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	return NewQiNiuOSSService(ctx, uploadManager, c.Bucket, c.Prefix, c.Url)
}
