package upload

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/now"
	"io"
	"sync"
	"time"
)

var (
	ErrOSSInitFail      = errors.New("OSS初始化失败，请联系管理员")
	ErrFileAlreadyExist = errors.New("文件已存在")
)

type ossUploader struct {
	_lock           sync.Mutex
	_init           bool
	_bucket         *oss.Bucket
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string
	saveFolder      string
}

func (o *ossUploader) UploadFile(reader io.Reader, filename string) (path string, err error) {
	o._initBucket()
	var uploadPath string
	if o._bucket == nil {
		return "", ErrOSSInitFail
	}
	uploadPath = fmt.Sprintf("%s/%s/%s", o.saveFolder, now.New(time.Now()).Format("2006/01/02"), filename)
	exist, err := o._bucket.IsObjectExist(uploadPath)
	if err != nil {
		return uploadPath, err
	}
	if exist {
		return uploadPath, ErrFileAlreadyExist
	}
	return uploadPath, o._bucket.PutObject(uploadPath, reader)
}

func (o *ossUploader) _initBucket() {
	// 单例锁
	if o._init {
		return
	}
	o._lock.Lock()
	defer o._lock.Unlock()
	if o._init {
		return
	}
	o._init = true

	client, err := oss.New(o.endpoint, o.accessKeyId, o.accessKeySecret)
	if err != nil {
		log.Default.Errorf("init oss fail with reason: %s", err.Error())
		return
	}
	o._bucket, err = client.Bucket(setting.Config.OSS.BucketName)
	if err != nil {
		log.Default.Errorf("init oss fail with reason: %s", err.Error())
	}
}
