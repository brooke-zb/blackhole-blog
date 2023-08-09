package util

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/now"
	"io"
	"time"
)

var bucket *oss.Bucket

var (
	ErrOSSInitFail      = errors.New("OSS初始化失败，请联系管理员")
	ErrFileAlreadyExist = errors.New("文件已存在")
)

func initOSS() {
	client, err := oss.New(setting.Config.OSS.Endpoint, setting.Config.OSS.AccessKeyId, setting.Config.OSS.AccessKeySecret)
	if err != nil {
		log.Default.Errorf("init oss fail with reason: %s", err.Error())
		return
	}
	bucket, err = client.Bucket(setting.Config.OSS.BucketName)
	if err != nil {
		log.Default.Errorf("init oss fail with reason: %s", err.Error())
	}
}

func UploadFile(reader io.Reader, filename string) (string, error) {
	var uploadPath string
	if bucket == nil {
		return "", ErrOSSInitFail
	}
	uploadPath = fmt.Sprintf("%s/%s/%s", setting.Config.OSS.SaveFolder, now.New(time.Now()).Format("2006/01/02"), filename)
	exist, err := bucket.IsObjectExist(setting.Config.OSS.SaveFolder + uploadPath + filename)
	if err != nil {
		return uploadPath, err
	}
	if exist {
		return uploadPath, ErrFileAlreadyExist
	}
	return uploadPath, bucket.PutObject(setting.Config.OSS.SaveFolder+uploadPath+filename, reader)
}
