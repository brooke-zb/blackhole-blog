package main

import (
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/upload"
	"blackhole-blog/pkg/util"
	"blackhole-blog/router"
	"blackhole-blog/task"
	"fmt"
)

func init() {
	setting.Setup()
	log.Setup()
	util.Setup()
	dao.Setup()
	task.Setup()
	upload.Setup()
}

func main() {
	r := router.InitRouter()
	r.Run(fmt.Sprintf("%s:%d", setting.Config.Server.Host, setting.Config.Server.Port))
}
