package ossx

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
)

type ALiYunService struct {
	logx.Logger
	client     *oss.Client
	ctx        context.Context
	bucketName string
	region     string
	Prefix     string
	endpoint   string
}

var _ Service = (*ALiYunService)(nil)

func NewALiYunOSSService(ctx context.Context, client *oss.Client, bucketName, region, prefix, endpoint string) *ALiYunService {
	return &ALiYunService{
		Logger:     logx.WithContext(ctx),
		client:     client,
		ctx:        ctx,
		bucketName: bucketName,
		region:     region,
		Prefix:     prefix,
		endpoint:   endpoint,
	}
}

func (s *ALiYunService) UploadFile(file *multipart.FileHeader) (string, error) {
	key := fmt.Sprintf("%s/%s", s.Prefix, file.Filename)
	reader, err := file.Open()
	if err != nil {
		s.Logger.Error(err)
		return "", FILE_READ_FAIL
	}
	defer reader.Close()
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(s.bucketName), // 存储空间名称
		Key:    oss.Ptr(key),          // 对象名称
		Body:   reader,                // 要上传的字符串内容
	}
	_, err = s.client.PutObject(s.ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return "", UPLOUD_FILE_FAIL
	}
	return fmt.Sprintf("https://%s.%s/%s", s.bucketName, s.endpoint, key), nil
}
