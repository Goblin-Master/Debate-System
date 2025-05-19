package ossx

import (
	"context"
	"errors"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	qiniuCredentials "github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"mime/multipart"
)

var (
	FILE_READ_FAIL   = errors.New("文件读取失败")
	UPLOUD_FILE_FAIL = errors.New("上传文件失败")
)

type Service interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func InitSSOService(ctx context.Context, c ALiYun) Service {
	return InitALiYunOSS(ctx, c)
}

type QiNiu struct {
	Enable    bool   `json:"enable"`
	AccessKey string `json:"accessKey"`
	Bucket    string `json:"bucket"`
	SecretKey string `json:"secretKey"`
	Url       string `json:"url"`
	Region    string `json:"region"`
	Prefix    string `json:"prefix"`
}

func InitQiNiuOSS(ctx context.Context, c QiNiu) Service {
	mac := qiniuCredentials.NewCredentials(c.AccessKey, c.SecretKey)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	return NewQiNiuOSSService(ctx, uploadManager, c.Bucket, c.Prefix, c.Url)
}

type ALiYun struct {
	Enable    bool   `json:"enable"`
	AccessID  string `json:"accessID"`
	Bucket    string `json:"bucket"`
	AccessKey string `json:"accessKey"`
	Region    string `json:"region"`
	Prefix    string `json:"prefix"`
	Endpoint  string `json:"endpoint"`
}

func InitALiYunOSS(ctx context.Context, c ALiYun) Service {
	// 使用NewStaticCredentialsProvider方法直接设置AK和SK
	provider := credentials.NewStaticCredentialsProvider(c.AccessID, c.AccessKey)
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(c.Region).
		WithEndpoint(c.Endpoint)
	// 创建OSS客户端
	client := oss.NewClient(cfg)
	return NewALiYunOSSService(ctx, client, c.Bucket, c.Region, c.Prefix, c.Endpoint)
}
