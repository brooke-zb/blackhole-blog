package task

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"github.com/go-co-op/gocron"
	"time"
)

var scheduler *gocron.Scheduler

func Setup() {
	scheduler = gocron.NewScheduler(time.Local)
	_, err := scheduler.CronWithSeconds(setting.Config.Task.Cron.PersistArticleReadCount).Do(persistArticleReadCountTask)
	if err != nil {
		log.Default.Errorf("init persistArticleReadCountTask fail with reason: %s", err.Error())
	}

	scheduler.StartAsync()
}
