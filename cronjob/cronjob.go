package cronjob

import (
	"github.com/penril0326/shorturl/model"
	"github.com/robfig/cron/v3"
)

func init() {
	c = cron.New()
	c.AddFunc("@weekly", func() {
		model.DeleteExpired()
	})
}

func Start() {
	c.Start()
}

func Stop() {
	c.Stop()
}
