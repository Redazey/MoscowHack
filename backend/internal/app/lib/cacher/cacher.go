package cacher

import (
	"fmt"
	"moscowhack/pkg/cache"

	"github.com/robfig/cron/v3"
)

func cacheUpdate() {
	cache.DeleteEX("*")
}

func Init(interval string) {
	intervalStr := fmt.Sprintf("%s * * * *", interval)

	c := cron.New()
	_, err := c.AddFunc(intervalStr, cacheUpdate)
	if err != nil {
		return
	}

	c.Start()
}
