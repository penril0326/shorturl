package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/penril0326/shorturl/cache"
	"github.com/penril0326/shorturl/controller/middleware"
	"github.com/penril0326/shorturl/controller/webapi"
	"github.com/penril0326/shorturl/cronjob"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	cronjob.Start()
}

func main() {
	defer cronjob.Stop()
	defer cache.DeleteAll()

	r := gin.Default()

	r.POST("/api/v1/urls", middleware.PostRequestLimit, webapi.CreateShort)
	r.DELETE("/api/v1/urls/:url_id", webapi.DeleteUrl)
	r.GET("/:url_id", middleware.GetRequestLimit, webapi.Redirect)

	go r.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown...")
}
