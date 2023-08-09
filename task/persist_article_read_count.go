package task

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"fmt"
	"strconv"
)

func persistArticleReadCountTask() {
	log.Default.Info("PersistArticleReadCountTask: task start")

	successCount := 0
	failCount := 0
	defer func() {
		log.Default.Infof("PersistArticleReadCountTask: task finish, success: %d, fail: %d", successCount, failCount)
	}()
	keys, err := util.Redis.Keys(fmt.Sprintf("%s*", setting.ArticleReadCountPrefix))
	if err != nil {
		log.Default.Errorf("PersistArticleReadCountTask: get keys fail with reason: %s", err.Error())
		return
	}
	for _, key := range keys {
		func() {
			// get article id
			aid, err := strconv.ParseUint(key[len(setting.ArticleReadCountPrefix):], 10, 64)
			var count int64 = 0
			defer func() {
				if err := recover(); err != nil {
					failCount++
					log.Default.Errorf("PersistArticleReadCountTask: %s", err)
				} else {
					successCount++
					log.Default.Infof("PersistArticleReadCountTask: persist success, article id: %d, readcount increment: %d", aid, count)
				}
			}()

			if err != nil {
				panic(fmt.Errorf("parse article id fail with reason: %s", err.Error()))
			}

			// get read count increment
			count, err = util.Redis.PFCount(key)
			if err != nil {
				panic(fmt.Errorf("get key %s fail with reason: %s", key, err.Error()))
			}

			// update read count
			service.Article.UpdateReadCount(aid, count)

			// delete key
			err = util.Redis.Del(key)
			if err != nil {
				panic(fmt.Errorf("delete key %s fail with reason: %s", key, err.Error()))
			}
		}()
	}
}
