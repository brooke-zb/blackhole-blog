package util

import "blackhole-blog/pkg/log"

func Setup() {
	InitRedis()
	initJwt()
	transErr := InitTrans("zh")
	if transErr != nil {
		log.Default.Error(transErr.Error())
	}
}
