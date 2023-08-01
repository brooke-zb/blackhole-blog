package main

import (
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/router"
	"fmt"
)

func init() {
	setting.Setup()
	log.Setup()
	dao.Setup()
	util.Setup()
}

func main() {
	router := router.InitRouter()
	router.Run(fmt.Sprintf("%s:%d", setting.Config.Server.Host, setting.Config.Server.Port))
}
