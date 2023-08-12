package upload

import (
	"blackhole-blog/pkg/setting"
	"io"
)

type uploader interface {
	UploadFile(reader io.Reader, filename string) (path string, err error)
}

var Uploader uploader

func Setup() {
	// use Aliyun OSS
	Uploader = &ossUploader{
		endpoint:        setting.Config.OSS.Endpoint,
		accessKeyId:     setting.Config.OSS.AccessKeyId,
		accessKeySecret: setting.Config.OSS.AccessKeySecret,
		bucketName:      setting.Config.OSS.BucketName,
		saveFolder:      setting.Config.OSS.SaveFolder,
	}
}
