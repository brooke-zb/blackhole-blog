package util

func Setup() {
	initRedis()
	initJwt()
	initTrans("zh")
	initWordFilter()
	initIdGenerator()
	initOSS()
}
