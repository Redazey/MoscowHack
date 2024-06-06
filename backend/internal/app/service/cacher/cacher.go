package cacher

import (
	"fmt"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"

	"github.com/robfig/cron/v3"
)

func cacheUpdate() {
	cacheTables, err := cache.ReadCache("*")
	if err != nil {
		return
	}

	for table := range cacheTables {
		cacheMap, err := cache.ReadCache(table)
		if err != nil {
			continue
		}

		db.PullData(table, cacheMap)
	}
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
