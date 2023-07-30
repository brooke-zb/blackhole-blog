package main

import (
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/routes"
	"fmt"
)

func init() {
	setting.Setup()
	dao.Setup()
}

func main() {
	router := routes.InitRouter()
	router.Run(fmt.Sprintf("%s:%d", setting.Config.Server.Host, setting.Config.Server.Port))
}
