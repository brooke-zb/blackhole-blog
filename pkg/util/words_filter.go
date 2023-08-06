package util

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"github.com/importcjj/sensitive"
)

var WordsFilter *sensitive.Filter

func initWordFilter() {
	WordsFilter = sensitive.New()
	if setting.Config.WordsFilter.WordsPath != nil {
		err := WordsFilter.LoadWordDict(*setting.Config.WordsFilter.WordsPath)
		if err != nil {
			log.Default.Error("init words filter fail with reason: " + err.Error())
		}
	}
}
