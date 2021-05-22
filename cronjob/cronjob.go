package cronjob

import (
	"log"

	"github.com/penril0326/shorturl/model"
	"github.com/robfig/cron/v3"
)

func init() {
	c = cron.New()
	c.AddFunc("@weekly", func() {
		if err := model.DeleteExpired(); err != nil {
			log.Println("Weekly delete expire data failed. Error: ", err.Error())
		}
	})
}

func Start() {
	c.Start()
}

func Stop() {
	c.Stop()
}
